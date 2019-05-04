package core

import (
	"os"
	"runtime"
)

type SystemInfo struct {
	OS       string
	ARCH     string
	NumCPU   int
	HostName string
}

func (systemInfo *SystemInfo) InitSystemInfo() {
	systemInfo.OS = runtime.GOOS
	systemInfo.ARCH = runtime.GOARCH
	systemInfo.NumCPU = runtime.NumCPU()
	systemInfo.HostName, _ = os.Hostname()
}
