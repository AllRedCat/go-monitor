package models

type CPUInfo struct {
	CPUI        int32   `json:"cpu"`
	SteppingI   int32   `json:"stepping"`
	PhysicalIdI string  `json:"physicalId`
	CoreIdI     string  `json:"coreId"`
	CoresI      int32   `json:"cores"`
	ModelNameI  string  `json:"modelName"`
	MhzI        float64 `json:"mhz"`
	ChaceSizeI  int32   `json:"cacheSize"`
	MicrocodeI  string  `json:"microencode"`
}
