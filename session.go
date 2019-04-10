package polyanalyst6api

import (
	"encoding/json"
	"errors"
	"fmt"
	"runtime"
	"strings"

	"github.com/Megaputer/polyanalyst6api-go/parameters/dataset"
	"github.com/Megaputer/polyanalyst6api-go/parameters/project"
	"github.com/Megaputer/polyanalyst6api-go/parameters/scheduler"
	"github.com/Megaputer/polyanalyst6api-go/responses"

	"github.com/Megaputer/polyanalyst6api-go/objects"
	"github.com/Megaputer/polyanalyst6api-go/parameters"
)

// Session is used to interact with the API
type Session struct {
	SID        string
	Server     *Server
	apiVersion string
}

// ProjectNodes returns the list of project nodes `/project/nodes`
func (s Session) ProjectNodes(uuid string) ([]objects.Node, error) {
	var nodes []objects.Node

	param := project.Nodes{PrjUUID: uuid}
	nodesData, err := s.request("GET", "/project/nodes", param.ToFullParams())
	if err != nil {
		return nodes, err
	}

	var nodesResp responses.Nodes
	err = json.Unmarshal(nodesData, &nodesResp)
	if err != nil {
		return nodes, err
	}

	return nodesResp.Nodes, nil
}

// ProjectExecutionStatistics returns execution statistics for specific project `/project/execution-statistics`
func (s Session) ProjectExecutionStatistics(uuid string) (responses.ProjectExecutionStatistics, error) {
	var statsResp responses.ProjectExecutionStatistics

	param := project.ExecutionStatistics{PrjUUID: uuid}
	nodesData, err := s.request("GET", "/project/execution-statistics", param.ToFullParams())
	if err != nil {
		return statsResp, err
	}

	err = json.Unmarshal(nodesData, &statsResp)
	if err != nil {
		return statsResp, err
	}

	return statsResp, nil
}

// ProjectExecute starts project execution `/project/execution-statistics`
func (s Session) ProjectExecute(params project.Execute) error {
	_, err := s.request("POST", "/project/execute", params.ToFullParams())
	return err
}

// ProjectGlobalAbort stops project execution: `/project/global-abort`
func (s Session) ProjectGlobalAbort(uuid string) error {
	params := project.GlobalAbort{PrjUUID: uuid}
	_, err := s.request("POST", "/project/global-abort", params.ToFullParams())
	return err
}

// ProjectSave initiates project saving: `/project/save`
func (s Session) ProjectSave(uuid string) error {
	params := project.Save{PrjUUID: uuid}
	_, err := s.request("POST", "/project/save", params.ToFullParams())
	return err
}

// ProjectUnload initiates project unloading: `/project/unload`
func (s Session) ProjectUnload(uuid string) error {
	params := project.Unload{PrjUUID: uuid}
	_, err := s.request("POST", "/project/unload", params.ToFullParams())
	return err
}

// ProjectRepair initiates project repairing: `/project/repair`
func (s Session) ProjectRepair(uuid string) error {
	params := project.Repair{PrjUUID: uuid}
	_, err := s.request("POST", "/project/repair", params.ToFullParams())
	return err
}

// ProjectDelete initiates project repairing: `/project/delete`
func (s Session) ProjectDelete(uuid string, forceUnload bool) error {
	params := project.Delete{PrjUUID: uuid, ForceUnload: forceUnload}
	_, err := s.request("POST", "/project/delete", params.ToFullParams())
	return err
}

// DatasetPreview returns first 1k records of the dataset: `/dataset/preview`
func (s Session) DatasetPreview(prjUUID string, name string, nodeType string) ([][]byte, error) {
	res := [][]byte{}

	if nodeType != "DataSource" && nodeType != "Dataset" {
		return res, fmt.Errorf("invalid node type: '%s' (only 'DataSource' and 'Dataset' types are allowed)", nodeType)
	}

	params := dataset.Preview{PrjUUID: prjUUID, Name: name, Type: nodeType}
	resp, err := s.request("GET", "/dataset/preview", params.ToFullParams())
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(resp, &res)

	return res, err
}

// SchedulerRunTask starts the task with passed ID
func (s Session) SchedulerRunTask(taskID uint) error {
	params := scheduler.RunTask{TaskID: taskID}
	_, err := s.request("POST", "/scheduler/run-task", params.ToFullParams())
	return err
}

// ProjectTasks returns the list of project tasks: `/project/tasks`
func (s Session) ProjectTasks(uuid string) ([]objects.ProjectTaskInfo, error) {
	var tasks []objects.ProjectTaskInfo

	param := project.Tasks{PrjUUID: uuid}
	tasksData, err := s.request("GET", "/project/tasks", param.ToFullParams())
	if err != nil {
		return tasks, err
	}

	var tasksResp responses.ProjectTasks
	err = json.Unmarshal(tasksData, &tasksResp)
	if err != nil {
		return tasks, err
	}

	return tasksResp, nil
}

// request is used for making requests to the API
func (s Session) request(reqType string, path string, params parameters.Full) ([]byte, error) {
	var data []byte

	if !checkPathSupported(s.apiVersion, path) {
		var methodName string

		pc, _, _, ok := runtime.Caller(1)
		details := runtime.FuncForPC(pc)
		if ok && details != nil {
			methodName = strings.Split(details.Name(), "Session.")[1]
		} else {
			methodName = "n/a"
		}

		versions := pathSupportedIn(path)
		vstr := "none"
		if len(versions) > 0 {
			vstr = strings.Join(versions, ", ")
		}
		return data, errors.New("`" + methodName + "` call is not supported in the API version " + s.apiVersion + "; versions that support: " + vstr)
	}

	fullURL := s.Server.BaseURL() + "/v" + s.apiVersion + path
	r, err := createRequest(fullURL, reqType, params)
	if err != nil {
		return data, err
	}
	r.UseSession(&s)

	return r.Perform()
}
