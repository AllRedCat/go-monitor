package services

import "github.com/shirou/gopsutil/v4/docker"

func GetContainers() ([]docker.CgroupDockerStat, error) {
	d, err := docker.GetDockerStat()
	if err != nil {
		return nil, err
	}

	return d, nil
}