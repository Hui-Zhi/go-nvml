package gonvml

import (
	"testing"
	"github.com/hui-zhi/go-nvml"
)

func TestNVMLInit(t *testing.T) {
	NVMLInit()
	t.Success()
}

func TestNVMLShutdown(t *testing.T) {
	NVMLShutdown()
	t.Success()
}


