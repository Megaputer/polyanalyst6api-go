package project

import (
	"encoding/json"
	"net/url"

	"github.com/gluk-skywalker/polyanalyst6api-go/parameters"
)

// Delete is the structure for Delete: `/project/delete`
type Delete struct {
	PrjUUID     string `json:"prjUUID"`
	ForceUnload bool   `json:"forceUnload"`
}

// ToFullParams turns the Delete to full request params
func (p Delete) ToFullParams() parameters.Full {
	return parameters.Full{URLParams: p.toURLParams(), BodyParams: p.toJSON()}
}

func (p Delete) toURLParams() url.Values {
	return nil
}

func (p Delete) toJSON() []byte {
	jsonParams, _ := json.Marshal(p)
	return jsonParams
}
