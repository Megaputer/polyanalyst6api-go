package parameters

import (
	"encoding/json"
	"net/url"
)

// ProjectDelete is the structure for ProjectDelete: `/project/delete`
type ProjectDelete struct {
	PrjUUID     string `json:"prjUUID"`
	ForceUnload bool   `json:"forceUnload"`
}

// ToFullParams turns the ProjectDelete to full request params
func (p ProjectDelete) ToFullParams() FullParams {
	return FullParams{URLParams: p.toURLParams(), BodyParams: p.toJSON()}
}

func (p ProjectDelete) toURLParams() url.Values {
	return nil
}

func (p ProjectDelete) toJSON() []byte {
	jsonParams, _ := json.Marshal(p)
	return jsonParams
}
