package parameters

import (
	"net/url"
)

// ProjectNodes is the structure for ProjectNodes: `/project/nodes`
type ProjectNodes struct {
	PrjUUID string `json:"prjUUID"`
}

// ToFullParams turns the ProjectExecute to full request params
func (p ProjectNodes) ToFullParams() FullParams {
	return FullParams{URLParams: p.toURLParams(), BodyParams: p.toJSON()}
}

func (p ProjectNodes) toURLParams() url.Values {
	params := url.Values{}
	params.Set("prjUUID", p.PrjUUID)
	return params
}

func (p ProjectNodes) toJSON() []byte {
	return nil
}
