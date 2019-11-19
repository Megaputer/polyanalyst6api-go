package userfolder

import (
	"encoding/json"
	"net/url"

	"github.com/Megaputer/polyanalyst6api-go/parameters"
)

// Create is parameters structure for create folder request: `/project/delete`
type Create struct {
	Path string `json:"path"`
	Name string `json:"name"`
}

// ToFullParams turns the Create parameters into request params
func (p Create) ToFullParams() parameters.Full {
	return parameters.Full{URLParams: p.toURLParams(), BodyParams: p.toJSON()}
}

func (p Create) toURLParams() url.Values {
	return nil
}

func (p Create) toJSON() []byte {
	jsonParams, _ := json.Marshal(p)
	return jsonParams
}
