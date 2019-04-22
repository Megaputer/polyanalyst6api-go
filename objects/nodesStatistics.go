package objects

// NodesStatistics is the structure for overall nodes statistics
type NodesStatistics struct {
	EmtpyNodesCount          int64 `json:"emtpyNodesCount"`
	SynchronizedNodesCount   int64 `json:"synchronizedNodesCount"`
	UnsynchronizedNodesCount int64 `json:"unsynchronizedNodesCount"`
}
