package objects

// NodesStatistics is the structure for overall nodes statistics
type NodesStatistics struct {
	EmtpyNodesCount          uint `json:"emtpyNodesCount"`
	SynchronizedNodesCount   uint `json:"synchronizedNodesCount"`
	UnsynchronizedNodesCount uint `json:"unsynchronizedNodesCount"`
}
