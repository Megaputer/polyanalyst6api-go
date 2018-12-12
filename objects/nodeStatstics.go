package objects

// NodeStatistics is the structure for node statistics
type NodeStatistics struct {
	ID                uint    `json:"id"`
	Type              string  `json:"type"`
	Name              string  `json:"name"`
	Status            string  `json:"status"`
	ErrMsg            string  `json:"errMsg"`
	StartTime         float32 `json:"startTime"`
	EndTime           float32 `json:"endTime"`
	Duration          float32 `json:"duration"`
	FreeMemoryInitial float32 `json:"freeMemoryInitial"`
	FreeMemoryFinal   float32 `json:"freeMemoryFinal"`
	FreeDiscInitial   float32 `json:"freeDiscInitial"`
	FreeDiscFinal     float32 `json:"freeDiscFinal"`
}
