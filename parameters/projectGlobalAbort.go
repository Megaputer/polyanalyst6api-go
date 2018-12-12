package parameters

import (
	"encoding/json"
	"net/url"
)

// ProjectGlobalAbort is the structure for ProjectGlobalAbort: `/project/global-abort`
type ProjectGlobalAbort struct {
	PrjUUID string `json:"prjUUID"`
}

// ToFullParams turns the ProjectGlobalAbort to full request params
func (p ProjectGlobalAbort) ToFullParams() FullParams {
	return FullParams{URLParams: p.toURLParams(), BodyParams: p.toJSON()}
}

func (p ProjectGlobalAbort) toURLParams() url.Values {
	return nil
}

func (p ProjectGlobalAbort) toJSON() []byte {
	jsonParams, _ := json.Marshal(p)
	return jsonParams
}
