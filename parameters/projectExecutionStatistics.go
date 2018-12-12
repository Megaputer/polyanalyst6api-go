package parameters

import (
	"net/url"
)

// ProjectExecutionStatistics is the structure for ProjectExecutionStatistics: `/project/execution-statistics`
type ProjectExecutionStatistics struct {
	PrjUUID string `json:"prjUUID"`
}

// ToFullParams turns the ProjectExecute to full request params
func (p ProjectExecutionStatistics) ToFullParams() FullParams {
	return FullParams{URLParams: p.toURLParams(), BodyParams: p.toJSON()}
}

func (p ProjectExecutionStatistics) toURLParams() url.Values {
	params := url.Values{}
	params.Set("prjUUID", p.PrjUUID)
	return params
}

func (p ProjectExecutionStatistics) toJSON() []byte {
	return nil
}
