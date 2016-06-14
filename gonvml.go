// Reference: https://github.com/nvidia/nvidia-docker

package gonvml


// #cgo LDFLAGS: -lnvidia-ml -L /usr/src/gdk/nvml/lib/
// #include <nvml.h>
import "C"

import (
	"fmt"
)


const (
	driverBufferSize = C.NVML_SYSTEM_DRIVER_VERSION_BUFFER_SIZE
)

func NVMLInit(){
	C.nvmlInit()
}

func NVMLShutdown(){
	C.nvmlShutdown()
}

func nvmlError(ret C.nvmlReturn_t) error {
	if ret == C.NVML_SUCCESS {
		return nil
	}

	err := C.GoString(C.nvmlErrorString(ret))

	return fmt.Errorf("NVML error: %v", err)
}

// In the current version, it's not used.
func GetDriverVersion() (string, error) {
	var driver [driverBufferSize]C.char
	ret := C.nvmlSystemGetDriverVersion(&driver[0], driverBufferSize)

	return C.GoString(&driver[0]), nvmlError(ret)
}

func GetDeviceCount() (uint, error) {
	var num C.uint

	err := nvmlError(C.nvmlDeviceGetCount(&num))

	return uint(num), err
}

func GetDevicePath(idx uint) (string, error) {
	var dev C.nvmlDevice_t
	var minor C.uint

	err := nvmlError(C.nvmlDeviceGetHandleByIndex(C.uint(idx), &dev))

	if err != nil {
		return "", err
	}

	err = nvmlError(C.nvmlDeviceGetMinorNumber(dev, &minor))
	path := fmt.Sprintf("/dev/nvidia%d", uint(minor))

	return path, err
}

