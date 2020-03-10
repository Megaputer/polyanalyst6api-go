package configureparameters

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/Megaputer/polyanalyst6api-go/parameters"
)

// ConfigureParameters is the structure for Preview: `/parameters/configure`
type ConfigureParameters struct {
	PrjUUID       string
	ObjID         int
	DeclareUnsync bool
	Parameters    interface {
		nodeType() string
		settings() interface{}
	}
}

// ToFullParams turns the ConfigureParameters to full request params
func (p ConfigureParameters) ToFullParams() parameters.Full {
	return parameters.Full{URLParams: p.toURLParams(), BodyParams: p.toJSON()}
}

func (p ConfigureParameters) toURLParams() url.Values {
	return url.Values{
		"prjUUID": {p.PrjUUID},
		"obj":     {strconv.Itoa(p.ObjID)},
	}
}

func (p ConfigureParameters) toJSON() []byte {
	var params struct {
		NodeType      string      `json:"type"`
		DeclareUnsync bool        `json:"declareUnsync"`
		Settings      interface{} `json:"settings"`
	}
	params.NodeType = p.Parameters.nodeType()
	params.Settings = p.Parameters.settings()
	params.DeclareUnsync = p.DeclareUnsync
	jsonParams, _ := json.Marshal(params)
	return jsonParams
}
