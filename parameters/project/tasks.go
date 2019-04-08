package project

import (
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
	params := url.Values{}
	params.Set("prjUUID", p.PrjUUID)
	return params
}

func (p Tasks) toJSON() []byte {
	return nil
}
