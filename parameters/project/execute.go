package project

import (
	"encoding/json"
	"net/url"

	"github.com/gluk-skywalker/polyanalyst6api-go/objects"
	"github.com/gluk-skywalker/polyanalyst6api-go/parameters"
)

// Execute is the structure for Execute: `/project/execute`
type Execute struct {
	PrjUUID string         `json:"prjUUID"`
	Nodes   []objects.Node `json:"nodes"`
}

// ToFullParams turns the Execute to full request params
func (p Execute) ToFullParams() parameters.Full {
	return parameters.Full{URLParams: p.toURLParams(), BodyParams: p.toJSON()}
}

func (p Execute) toURLParams() url.Values {
	return nil
}

func (p Execute) toJSON() []byte {
	jsonParams, _ := json.Marshal(p)
	return jsonParams
}
