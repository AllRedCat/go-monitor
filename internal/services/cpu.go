package services

import (
	"time"

	"github.com/shirou/gopsutil/v4/cpu"

	"monitor/internal/models"
)

func GetCpuUsage() (models.CPUUsage, error) {
	g, err := cpu.Percent(1000*time.Millisecond, false)
	if err != nil {
		return models.CPUUsage{}, err
	}
	percent := g[0]

	e, err := cpu.Percent(1000*time.Millisecond, true)
	if err != nil {
		return models.CPUUsage{}, err
	}

	data := models.CPUUsage{
		GeneralPercent: percent,
		EachPercent:    e,
	}

	return data, nil
}