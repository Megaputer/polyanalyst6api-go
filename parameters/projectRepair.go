package parameters

import (
	"encoding/json"
	"net/url"
)

// ProjectRepair is the structure for ProjectRepair: `/project/repair`
type ProjectRepair struct {
	PrjUUID string `json:"prjUUID"`
}

// ToFullParams turns the ProjectRepair to full request params
func (p ProjectRepair) ToFullParams() FullParams {
	return FullParams{URLParams: p.toURLParams(), BodyParams: p.toJSON()}
}

func (p ProjectRepair) toURLParams() url.Values {
	return nil
}

func (p ProjectRepair) toJSON() []byte {
	jsonParams, _ := json.Marshal(p)
	return jsonParams
}
