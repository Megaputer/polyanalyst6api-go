package objects

// ProjectTaskInfo represents struct for project task info
type ProjectTaskInfo struct {
	Name         string  `json:"name"`
	ObjID        int     `json:"objId"`
	Progress     float64 `json:"progress"`
	SubProgress  int     `json:"subProgress"`
	CurrentState string  `json:"currentState"`
	StartTime    int64   `json:"startTime"`
}
