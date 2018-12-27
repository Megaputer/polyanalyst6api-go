package project

import (
	"net/url"

	"github.com/Megaputer/polyanalyst6api-go/parameters"
)

// Nodes is the structure for Nodes: `/project/nodes`
type Nodes struct {
	PrjUUID string `json:"prjUUID"`
}

// ToFullParams turns the ProjectExecute to full request params
func (p Nodes) ToFullParams() parameters.Full {
	return parameters.Full{URLParams: p.toURLParams(), BodyParams: p.toJSON()}
}

func (p Nodes) toURLParams() url.Values {
	params := url.Values{}
	params.Set("prjUUID", p.PrjUUID)
	return params
}

func (p Nodes) toJSON() []byte {
	return nil
}
