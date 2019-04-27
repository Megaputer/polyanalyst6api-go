package objects

// ProjectTaskInfo represents struct for project task info
type ProjectTaskInfo struct {
	Name         string  `json:"name"`
	ObjID        int     `json:"objId"`
	Progress     float64 `json:"progress"`
	SubProgress  float64 `json:"subProgress"`
	CurrentState string  `json:"currentState"`
	StartTime    int     `json:"startTime"`
}
