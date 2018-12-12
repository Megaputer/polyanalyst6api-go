package parameters

import (
	"encoding/json"
	"net/url"

	"github.com/gluk-skywalker/polyanalyst6api-go/objects"
)

// ProjectExecute is the structure for ProjectExecute: `/project/execute`
type ProjectExecute struct {
	PrjUUID string         `json:"prjUUID"`
	Nodes   []objects.Node `json:"nodes"`
}

// ToFullParams turns the ProjectExecute to full request params
func (p ProjectExecute) ToFullParams() FullParams {
	return FullParams{URLParams: p.toURLParams(), BodyParams: p.toJSON()}
}

func (p ProjectExecute) toURLParams() url.Values {
	return nil
}

func (p ProjectExecute) toJSON() []byte {
	jsonParams, _ := json.Marshal(p)
	return jsonParams
}
