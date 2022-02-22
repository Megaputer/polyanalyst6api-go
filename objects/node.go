package objects

// Node is for node data
type Node struct {
	ID      int    `json:"id"`
	Type    string `json:"type"`
	Subtype string `json:"subtype"`
	Name    string `json:"name"`
	Status  string `json:"status"`
	ErrMsg  string `json:"errMsg"`
}
