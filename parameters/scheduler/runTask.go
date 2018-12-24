package scheduler

import (
	"encoding/json"
	"net/url"

	"github.com/gluk-skywalker/polyanalyst6api-go/parameters"
)

// RunTask is the struct for parameters of `/scheduler/run-task`
type RunTask struct {
	TaskID uint `json:"taskId"`
}

// ToFullParams turns the Unload to full request params
func (p RunTask) ToFullParams() parameters.Full {
	return parameters.Full{URLParams: p.toURLParams(), BodyParams: p.toJSON()}
}

func (p RunTask) toURLParams() url.Values {
	return nil
}

func (p RunTask) toJSON() []byte {
	jsonParams, _ := json.Marshal(p)
	return jsonParams
}
