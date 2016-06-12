// Need to have the below abilities:
// Reference: http://github.com/nvidia/nvidia-docker
 
package gonvml

// #cgo LDFLAGS: -ldl -Wl,--unresolved-symbols=ignore-in-object-files
// #include "nvml/gonvml.h"
import "C"

import (
	"fmt"
)

const (
	szDriver   = C.NVML_SYSTEM_DRIVER_VERSION_BUFFER_SIZE
	szModel    = C.NVML_DEVICE_NAME_BUFFER_SIZE
	szUUID     = C.NVML_DEVICE_UUID_BUFFER_SIZE
	szProcs    = 32
	szProcName = 64
)

var (
	ErrCPUAffinity        = errors.New("failed to retrieve CPU affinity")
	ErrUnsupportedP2PLink = errors.New("unsupported P2P link type")
)

type P2PLinkType uint

const (
	P2PLinkUnknown P2PLinkType = iota
	P2PLinkCrossCPU
	P2PLinkSameCPU
	P2PLinkHostBridge
	P2PLinkMultiSwitch
	P2PLinkSingleSwitch
	P2PLinkSameBoard
)

type P2PLink struct {
	BusID string
	Link  P2PLinkType
}

func (t P2PLinkType) String() string {
	switch t {
	case P2PLinkCrossCPU:
		return "Cross CPU socket"
	case P2PLinkSameCPU:
		return "Same CPU socket"
	case P2PLinkHostBridge:
		return "Host PCI bridge"
	case P2PLinkMultiSwitch:
		return "Multiple PCI switches"
	case P2PLinkSingleSwitch:
		return "Single PCI switch"
	case P2PLinkSameBoard:
		return "Same board"
	case P2PLinkUnknown:
	}
	return "???"
}

type ClockInfo struct{
	Cores 	uint
	Memory	uint
}

type PCIInfo struct {
	BusID			string
	BAR1			uint64
	Bandwidth	uint
}

type Device struct {
	Model				string
	UUID				string
	Path				string
	Power				uint
	CPUAffinity	uint
	PCI					PCIInfo
	Clocks			ClockInfo
	Topology		[]P2PLink
}

type UtilizationInfo struct {
	GPU     uint
	Memory  uint
	Encoder uint
	Decoder uint
}

type PCIThroughputInfo struct {
	RX uint
	TX uint
}

type PCIStatusInfo struct {
	BAR1Used   uint64
	Throughput PCIThroughputInfo
}

type ECCErrorsInfo struct {
	L1Cache uint64
	L2Cache uint64
	Global  uint64
}

type MemoryInfo struct {
	GlobalUsed uint64
	ECCErrors  ECCErrorsInfo
}

type ProcessInfo struct {
	PID        uint
	Name       string
	MemoryUsed uint64
}

type DeviceStatus struct {
	Power       uint
	Temperature uint
	Utilization UtilizationInfo
	Memory      MemoryInfo
	Clocks      ClockInfo
	PCI         PCIStatusInfo
	Processes   []ProcessInfo
}

func GetDeviceCount() (uint, error) {
	var n C.uint

	err := nvmlErr(C.nvmlDeviceGetCount(&n))
	return uint(n), err
}

func GetDevicePath(idx uint) (path string, err error) {
	var dev C.nvmlDevice_t
	var minor C.uint

	err = nvmlErr(C.nvmlDeviceGetHandleByIndex(C.uint(idx), &dev))
	if err != nil {
		return
	}
	err = nvmlErr(C.nvmlDeviceGetMinorNumber(dev, &minor))
	path = fmt.Sprintf("/dev/nvidia%d", uint(minor))
	return
}


