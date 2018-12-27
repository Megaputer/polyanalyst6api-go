package responses

import "github.com/Megaputer/polyanalyst6api-go/objects"

// ProjectExecutionStatistics is the struct for `/project/execution-statistics` request response
type ProjectExecutionStatistics struct {
	Nodes           []objects.NodeStatistics `json:"nodes"`
	NodesStatistics objects.NodesStatistics  `json:"nodesStatistics"`
}
