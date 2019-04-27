package objects

// NodeStatistics is the structure for node statistics
type NodeStatistics struct {
	ID                int   `json:"id"`
	Type              string  `json:"type"`
	Name              string  `json:"name"`
	Status            string  `json:"status"`
	ErrMsg            string  `json:"errMsg"`
	StartTime         float64 `json:"startTime"`
	EndTime           float64 `json:"endTime"`
	Duration          float64 `json:"duration"`
	DatasetRows       int   `json:"datasetRows"`
	DatasetCols       int   `json:"datasetCols"`
	FreeMemoryInitial float64 `json:"freeMemoryInitial"`
	FreeMemoryFinal   float64 `json:"freeMemoryFinal"`
	FreeDiscInitial   float64 `json:"freeDiscInitial"`
	FreeDiscFinal     float64 `json:"freeDiscFinal"`
}
