package userfolder

import (
	"encoding/json"
	"net/url"

	"github.com/Megaputer/polyanalyst6api-go/parameters"
)

// Delete is parameters structure for create folder request: `/project/delete`
type Delete struct {
	Path string `json:"path"`
	Name string `json:"name"`
}

// ToFullParams turns the Delete parameters into full request params
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
