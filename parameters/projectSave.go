package parameters

import (
	"encoding/json"
	"net/url"
)

// ProjectSave is the structure for ProjectSave: `/project/save`
type ProjectSave struct {
	PrjUUID string `json:"prjUUID"`
}

// ToFullParams turns the ProjectSave to full request params
func (p ProjectSave) ToFullParams() FullParams {
	return FullParams{URLParams: p.toURLParams(), BodyParams: p.toJSON()}
}

func (p ProjectSave) toURLParams() url.Values {
	return nil
}

func (p ProjectSave) toJSON() []byte {
	jsonParams, _ := json.Marshal(p)
	return jsonParams
}
