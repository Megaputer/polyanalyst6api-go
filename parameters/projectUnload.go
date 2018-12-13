package parameters

import (
	"encoding/json"
	"net/url"
)

// ProjectUnload is the structure for ProjectUnload: `/project/unload`
type ProjectUnload struct {
	PrjUUID string `json:"prjUUID"`
}

// ToFullParams turns the ProjectUnload to full request params
func (p ProjectUnload) ToFullParams() FullParams {
	return FullParams{URLParams: p.toURLParams(), BodyParams: p.toJSON()}
}

func (p ProjectUnload) toURLParams() url.Values {
	return nil
}

func (p ProjectUnload) toJSON() []byte {
	jsonParams, _ := json.Marshal(p)
	return jsonParams
}
