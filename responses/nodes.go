package responses

import "github.com/Megaputer/polyanalyst6api-go/objects"

// Nodes is the struct for `/project/nodes` request resoibse
type Nodes struct {
	Nodes []objects.Node `json:"nodes"`
}
