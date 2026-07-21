package server_Info

import (
	"go.uber.org/zap"
	"goskeleton/app/global/variable"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

// 获取CPU使用率
func GetCpuPercent() (float64, error) {
	percent, err := cpu.Percent(time.Second, true)
	if err != nil {
		variable.ZapLog.Error("GetCpuPercent 获取CPU使用率出错：", zap.Error(err))
		return 0, err
	}
	totalPercent := 0.0
	for _, p := range percent {
		totalPercent += p
	}
	return totalPercent, nil
}

// 获取内存使用率
func GetMemPercent() (float64, error) {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		variable.ZapLog.Error("GetMemPercent 获取内存使用率出错：", zap.Error(err))
		return 0, err
	}
	return memInfo.UsedPercent, nil
}
