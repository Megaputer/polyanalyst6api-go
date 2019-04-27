package objects

// NodesStatistics is the structure for overall nodes statistics
type NodesStatistics struct {
	EmtpyNodesCount          int `json:"emtpyNodesCount"`
	SynchronizedNodesCount   int `json:"synchronizedNodesCount"`
	UnsynchronizedNodesCount int `json:"unsynchronizedNodesCount"`
}
