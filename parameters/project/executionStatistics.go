package project

import (
	"net/url"

	"github.com/gluk-skywalker/polyanalyst6api-go/parameters"
)

// ExecutionStatistics is the structure for ExecutionStatistics: `/project/execution-statistics`
type ExecutionStatistics struct {
	PrjUUID string `json:"prjUUID"`
}

// ToFullParams turns the ProjectExecute to full request params
func (p ExecutionStatistics) ToFullParams() parameters.Full {
	return parameters.Full{URLParams: p.toURLParams(), BodyParams: p.toJSON()}
}

func (p ExecutionStatistics) toURLParams() url.Values {
	params := url.Values{}
	params.Set("prjUUID", p.PrjUUID)
	return params
}

func (p ExecutionStatistics) toJSON() []byte {
	return nil
}
