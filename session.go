package polyanalyst6api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gluk-skywalker/polyanalyst6api-go/parameters/dataset"
	"github.com/gluk-skywalker/polyanalyst6api-go/parameters/project"
	"github.com/gluk-skywalker/polyanalyst6api-go/parameters/scheduler"
	"github.com/gluk-skywalker/polyanalyst6api-go/responses"

	"github.com/gluk-skywalker/polyanalyst6api-go/objects"
	"github.com/gluk-skywalker/polyanalyst6api-go/parameters"
)

// Session is used to interact with the API
type Session struct {
	SID    string
	Server *Server
}

// ProjectNodes returns the list of project nodes `/project/nodes`
func (s Session) ProjectNodes(uuid string) ([]objects.Node, error) {
	var nodes []objects.Node

	param := project.Nodes{PrjUUID: uuid}
	nodesData, err := s.request("GET", "/project/nodes", param.ToFullParams())
	if err != nil {
		return nodes, errors.New(err.Error())
	}

	var nodesResp responses.Nodes
	err = json.Unmarshal(nodesData, &nodesResp)
	if err != nil {
		return nodes, errors.New(err.Error())
	}

	return nodesResp.Nodes, nil
}

// ProjectExecutionStatistics returns execution statistics for specific project `/project/execution-statistics`
func (s Session) ProjectExecutionStatistics(uuid string) (responses.ProjectExecutionStatistics, error) {
	var statsResp responses.ProjectExecutionStatistics

	param := project.ExecutionStatistics{PrjUUID: uuid}
	nodesData, err := s.request("GET", "/project/execution-statistics", param.ToFullParams())
	if err != nil {
		return statsResp, errors.New(err.Error())
	}

	err = json.Unmarshal(nodesData, &statsResp)
	if err != nil {
		return statsResp, errors.New(err.Error())
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
func (s Session) DatasetPreview(prjUUID string, name string, nodeType string) (string, error) {
	params := dataset.Preview{PrjUUID: prjUUID, Name: name, Type: nodeType}
	resp, err := s.request("GET", "/dataset/preview", params.ToFullParams())
	if err != nil {
		return "", err
	}

	return string(resp), err
}

// SchedulerRunTask starts the task with passed ID
func (s Session) SchedulerRunTask(taskID uint) error {
	params := scheduler.RunTask{TaskID: taskID}
	_, err := s.request("POST", "/scheduler/run-task", params.ToFullParams())
	return err
}

// request is used for making requests to the API
func (s Session) request(reqType string, path string, params parameters.Full) ([]byte, error) {
	var (
		err  error
		data []byte
	)

	// turning url paras to RFC 3986 compatible string
	urlParams := strings.Replace(params.URLParams.Encode(), "+", "%20", -1)

	url := s.Server.BaseURL() + path + "?" + urlParams
	req, err := http.NewRequest(reqType, url, bytes.NewBuffer(params.BodyParams))
	if err != nil {
		return data, errors.New("building request error: " + err.Error())
	}

	cookie := http.Cookie{Name: "sid", Value: s.SID}
	req.AddCookie(&cookie)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return data, errors.New("request execution error: " + err.Error())
	}
	defer resp.Body.Close()

	data, errBodyRead := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		msg := ""
		if errBodyRead != nil {
			msg = "*failed to retrieve*"
		} else {
			msg = string(data)
		}
		return data, fmt.Errorf("bad response status: %d. Error: %s", resp.StatusCode, msg)
	}

	if errBodyRead != nil {
		return data, errors.New("failed to read response")
	}

	return data, nil
}
