package polyanalyst6api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gluk-skywalker/polyanalyst6api-go/objects"
	"github.com/gluk-skywalker/polyanalyst6api-go/responses"
)

// Session is used to interact with the API
type Session struct {
	SID     string
	BaseURL string
}

// ProjectNodes returns the list of project nodes `/project/nodes`
func (s Session) ProjectNodes(uuid string) ([]objects.Node, error) {
	var nodes []objects.Node

	param := "prjUUID=" + uuid
	nodesData, err := s.request("GET", "/project/nodes", param)
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

// request is used for making requests to the API
func (s Session) request(reqType string, path string, params string) ([]byte, error) {
	var (
		err  error
		data []byte
	)

	url := s.BaseURL + path + "?" + params
	req, err := http.NewRequest(reqType, url, nil)
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

	if resp.StatusCode != 200 {
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
