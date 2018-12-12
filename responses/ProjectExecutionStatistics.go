package responses

import "github.com/gluk-skywalker/polyanalyst6api-go/objects"

// ProjectExecutionStatistics is the struct for `/project/execution-statistics` request resoibse
type ProjectExecutionStatistics struct {
	Nodes           []objects.NodeStatistics `json:"nodes"`
	NodesStatistics objects.NodesStatistics  `json:"nodesStatistics"`
}
