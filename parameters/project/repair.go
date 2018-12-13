package project

import (
	"encoding/json"
	"net/url"

	"github.com/gluk-skywalker/polyanalyst6api-go/parameters"
)

// Repair is the structure for Repair: `/project/repair`
type Repair struct {
	PrjUUID string `json:"prjUUID"`
}

// ToFullParams turns the Repair to full request params
func (p Repair) ToFullParams() parameters.Full {
	return parameters.Full{URLParams: p.toURLParams(), BodyParams: p.toJSON()}
}

func (p Repair) toURLParams() url.Values {
	return nil
}

func (p Repair) toJSON() []byte {
	jsonParams, _ := json.Marshal(p)
	return jsonParams
}
