package polyanalyst6api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"runtime"
	"strings"

	cfgparams "github.com/Megaputer/polyanalyst6api-go/parameters/configureParameters"
	"github.com/Megaputer/polyanalyst6api-go/parameters/dataset"
	"github.com/Megaputer/polyanalyst6api-go/parameters/project"
	"github.com/Megaputer/polyanalyst6api-go/parameters/scheduler"
	"github.com/Megaputer/polyanalyst6api-go/parameters/userfolder"
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
	err = json.Unmarshal(nodesData.Body, &nodesResp)
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

	err = json.Unmarshal(nodesData.Body, &statsResp)
	if err != nil {
		return statsResp, err
	}

	return statsResp, nil
}

// ProjectExecute starts project execution `/project/execution-statistics`
func (s Session) ProjectExecute(params project.Execute) (int, error) {
	result, err := s.request("POST", "/project/execute", params.ToFullParams())
	if err != nil {
		return 0, err
	}
	if result.Additions.ExecutionWaveID != nil && *result.Additions.ExecutionWaveID < 1 {
		return 0, fmt.Errorf("Bad execution wave ID")
	}
	return *result.Additions.ExecutionWaveID, nil
}

// ProjectIsRunning returns true if project is running, false otherwise
func (s Session) ProjectIsRunning(params project.IsRunning) (bool, error) {
	resp, err := s.request("GET", "/project/is-running", params.ToFullParams())
	if err != nil {
		return false, err
	}
	var res struct {
		Result int `json:"result"`
	}
	err = json.Unmarshal(resp.Body, &res)
	if err != nil {
		return false, err
	}

	return res.Result != 0, nil
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

	err = json.Unmarshal(resp.Body, &res)

	return res, err
}

// SchedulerRunTask starts the task with passed ID
func (s Session) SchedulerRunTask(taskID int) error {
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
	err = json.Unmarshal(tasksData.Body, &tasksResp)
	if err != nil {
		return tasks, err
	}

	return tasksResp, nil
}

// ParametersConfigure sets the parameters of Parmeters node
func (s Session) ParametersConfigure(param cfgparams.ConfigureParameters) error {
	_, err := s.request("POST", "/parameters/configure", param.ToFullParams())
	return err
}

// ServerInfo returns the list of project tasks: `/server/info`
func (s Session) ServerInfo() (responses.ServerInfo, error) {
	var info responses.ServerInfo

	tasksData, err := s.request("GET", "/server/info", parameters.Full{})
	if err != nil {
		return info, err
	}

	err = json.Unmarshal(tasksData.Body, &info)
	if err != nil {
		return info, err
	}

	return info, nil
}

// ParametersAvailableNodes returns the available nodes that configured via Parameters node: `/parameters/nodes`
func (s Session) ParametersAvailableNodes() ([]responses.ParametersAvailableNode, error) {
	var res []responses.ParametersAvailableNode

	data, err := s.request("GET", "/parameters/nodes", parameters.Full{})
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data.Body, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

// UserFolderCreate creates a folder in the user directory: `/folder/create`
func (s Session) UserFolderCreate(path string, name string) error {
	params := userfolder.Create{
		Path: path,
		Name: name,
	}
	_, err := s.request("POST", "/folder/create", params.ToFullParams())
	return err
}

// UserFolderDelete creates a folder in the user directory: `/folder/delete`
func (s Session) UserFolderDelete(path string, name string) error {
	params := userfolder.Delete{
		Path: path,
		Name: name,
	}
	_, err := s.request("POST", "/folder/delete", params.ToFullParams())
	return err
}

// Logout log out the user: `/logout`
func (s Session) Logout() error {
	params := parameters.Full{}
	_, err := s.request("GET", "/logout", params)
	return err
}

// request is used for making requests to the API
func (s Session) request(reqType string, path string, params parameters.Full) (RequestResult, error) {
	var result RequestResult

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
		return result, fmt.Errorf("'%s' call is not supported in the API version %s; versions that support: %s", methodName, s.apiVersion, vstr)
	}

	fullURL := s.Server.BaseURL() + "/v" + s.apiVersion + path
	r, err := createRequest(fullURL, reqType, params)
	if err != nil {
		return result, err
	}
	r.UseSession(&s)

	return r.Perform()
}

func (s Session) CustomRequest(reqType string, path string, body []byte, result any) error {
	fullURL := s.Server.Address() + path

	req, err := http.NewRequest(reqType, fullURL, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("request creation failed: %s", err)
	}

	cookie := http.Cookie{Name: "sid", Value: s.SID}
	req.AddCookie(&cookie)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("bad status code: %d", resp.StatusCode)
	}

	if result != nil {
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("body read failed: %s", err)
		}

		err = json.Unmarshal(data, result)
		if err != nil {
			return fmt.Errorf("failed to parse response body: %w", err)
		}
	}

	return nil
}
