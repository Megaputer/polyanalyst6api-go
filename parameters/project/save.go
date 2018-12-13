package project

import (
	"encoding/json"
	"net/url"

	"github.com/gluk-skywalker/polyanalyst6api-go/parameters"
)

// Save is the structure for Save: `/project/save`
type Save struct {
	PrjUUID string `json:"prjUUID"`
}

// ToFullParams turns the Save to full request params
func (p Save) ToFullParams() parameters.Full {
	return parameters.Full{URLParams: p.toURLParams(), BodyParams: p.toJSON()}
}

func (p Save) toURLParams() url.Values {
	return nil
}

func (p Save) toJSON() []byte {
	jsonParams, _ := json.Marshal(p)
	return jsonParams
}
