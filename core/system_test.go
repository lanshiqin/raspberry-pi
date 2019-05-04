package core

import (
	"strconv"
	"testing"
)

func TestSystemInfo_InitSystemInfo(t *testing.T) {
	systemInfo := SystemInfo{}
	systemInfo.InitSystemInfo()
	t.Log("系统类型：" + systemInfo.OS)
	t.Log("系统架构：" + systemInfo.ARCH)
	t.Log("CPU核心数：" + strconv.Itoa(systemInfo.NumCPU))
	t.Log("主机名称：" + systemInfo.HostName)

}
