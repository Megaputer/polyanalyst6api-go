package project

import (
	"encoding/json"
	"net/url"

	"github.com/Megaputer/polyanalyst6api-go/parameters"
)

// Unload is the structure for Unload: `/project/unload`
type Unload struct {
	PrjUUID string `json:"prjUUID"`
}

// ToFullParams turns the Unload to full request params
func (p Unload) ToFullParams() parameters.Full {
	return parameters.Full{URLParams: p.toURLParams(), BodyParams: p.toJSON()}
}

func (p Unload) toURLParams() url.Values {
	return nil
}

func (p Unload) toJSON() []byte {
	jsonParams, _ := json.Marshal(p)
	return jsonParams
}
