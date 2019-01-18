package project

import (
	"encoding/json"
	"net/url"

	"github.com/Megaputer/polyanalyst6api-go/parameters"
)

// Tasks is the structure for Tasks: `/project/tasks`
type Tasks struct {
	PrjUUID string `json:"prjUUID"`
}

// ToFullParams turns the Tasks to full request params
func (p Tasks) ToFullParams() parameters.Full {
	return parameters.Full{URLParams: p.toURLParams(), BodyParams: p.toJSON()}
}

func (p Tasks) toURLParams() url.Values {
	return nil
}

func (p Tasks) toJSON() []byte {
	jsonParams, _ := json.Marshal(p)
	return jsonParams
}
