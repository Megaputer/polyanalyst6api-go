package dataset

import (
	"net/url"

	"github.com/gluk-skywalker/polyanalyst6api-go/parameters"
)

// Preview is the structure for Preview: `/project/preview`
type Preview struct {
	PrjUUID string `json:"prjUUID"`
	Name    string `json:"name"`
	Type    string `json:"type"`
}

// ToFullParams turns the Execute to full request params
func (p Preview) ToFullParams() parameters.Full {
	return parameters.Full{URLParams: p.toURLParams(), BodyParams: p.toJSON()}
}

func (p Preview) toURLParams() url.Values {
	return url.Values{
		"prjUUID": {p.PrjUUID},
		"name":    {p.Name},
		"type":    {p.Type},
	}
}

func (p Preview) toJSON() []byte {
	return nil
}
