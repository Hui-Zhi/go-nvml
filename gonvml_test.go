package gonvml_test

import (
	"io/ioutil"
	"testing"

	"github.com/hui-zhi/gonvml"
)

func TestGetNvidiaGPUInfo(t *testing.T) {
	gpuInfo := gonvml.GetNvidiaGPUInfo()
}

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}

