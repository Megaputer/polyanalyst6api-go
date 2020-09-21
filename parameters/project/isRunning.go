package project

import (
	"net/url"
	"strconv"

	"github.com/Megaputer/polyanalyst6api-go/parameters"
)

// IsRunning is the structure for Execute: `/project/is-running`
type IsRunning struct {
	PrjUUID         string `json:"prjUUID"`
	ExecutionWaveID int    `json:"executionWave"`
}

// ToFullParams turns the Execute to full request params
func (p IsRunning) ToFullParams() parameters.Full {
	return parameters.Full{URLParams: p.toURLParams(), BodyParams: p.toJSON()}
}

func (p IsRunning) toURLParams() url.Values {
	params := url.Values{}
	params.Set("prjUUID", p.PrjUUID)
	params.Set("executionWave", strconv.Itoa(p.ExecutionWaveID))
	return params
}

func (p IsRunning) toJSON() []byte {
	return nil
}
