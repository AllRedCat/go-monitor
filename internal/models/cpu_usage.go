package models

type CPUUsage struct {
	GeneralPercent float64 `json:"general_percent"`
	EachPercent []float64 `json:"each_percent"`
}