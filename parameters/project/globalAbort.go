package project

import (
	"encoding/json"
	"net/url"

	"github.com/gluk-skywalker/polyanalyst6api-go/parameters"
)

// GlobalAbort is the structure for GlobalAbort: `/project/global-abort`
type GlobalAbort struct {
	PrjUUID string `json:"prjUUID"`
}

// ToFullParams turns the GlobalAbort to full request params
func (p GlobalAbort) ToFullParams() parameters.Full {
	return parameters.Full{URLParams: p.toURLParams(), BodyParams: p.toJSON()}
}

func (p GlobalAbort) toURLParams() url.Values {
	return nil
}

func (p GlobalAbort) toJSON() []byte {
	jsonParams, _ := json.Marshal(p)
	return jsonParams
}
