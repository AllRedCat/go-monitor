package services

import (
	"monitor/internal/models"

	"github.com/shirou/gopsutil/v4/cpu"
)

func GetMachineInfo() ([]models.CPUInfo, error) {
	cpus, err := cpu.Info()
	if err != nil {
		return nil, err
	}

	var infos []models.CPUInfo
	for _, i := range cpus {
		infos = append(infos, models.CPUInfo{
			CPUI:        i.CPU,
			SteppingI:   i.Stepping,
			PhysicalIdI: i.PhysicalID,
			CoreIdI:     i.CoreID,
			CoresI:      i.Cores,
			ModelNameI:  i.ModelName,
			MhzI:        i.Mhz,
			ChaceSizeI:  i.CacheSize,
			MicrocodeI:  i.Microcode,
		})
	}

	return infos, nil
}
