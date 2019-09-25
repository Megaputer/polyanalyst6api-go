package objects

// NodesStatistics is the structure for overall nodes statistics
type NodesStatistics struct {
	EmptyNodesCount          int `json:"emptyNodesCount"`
	SynchronizedNodesCount   int `json:"synchronizedNodesCount"`
	UnsynchronizedNodesCount int `json:"unsynchronizedNodesCount"`
}
