package polyanalyst6api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/Megaputer/polyanalyst6api-go/parameters"
)

type request struct {
	httpReq *http.Request
	// session *Session
	// path    string
	// reqType string
	// params  parameters.Full
}

type serverErrorData struct {
	Content ServerError `json:"error"`
}

func createRequest(path string, reqType string, params parameters.Full) (request, error) {
	var (
		retReq request
		err    error
	)

	// turning url paras to RFC 3986 compatible string
	urlParams := strings.Replace(params.URLParams.Encode(), "+", "%20", -1)

	url := path + "?" + urlParams

	req, err := http.NewRequest(reqType, url, bytes.NewBuffer(params.BodyParams))
	if err != nil {
		return retReq, fmt.Errorf("building request error: %s", err)
	}

	return request{httpReq: req}, nil
}

func (r *request) UseSession(s *Session) {
	cookie := http.Cookie{Name: "sid", Value: s.SID}
	r.httpReq.AddCookie(&cookie)
}

func (r request) Perform() ([]byte, error) {
	var (
		err  error
		data []byte
	)

	client := &http.Client{}

	resp, err := client.Do(r.httpReq)
	if err != nil {
		return data, fmt.Errorf("request execution error: %s", err)
	}
	defer closeBody(resp)

	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return data, fmt.Errorf("failed to read response body: %s", err)
	}

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		msg := string(data)
		var serverErr ServerError
		if isMain() {
			serverErr.Code = resp.StatusCode
			serverErr.Title = msg
			return data, serverErr
		}

		var errorData serverErrorData
		err = json.Unmarshal(data, &errorData)
		if err != nil {
			return data, fmt.Errorf("failed to parse server error [%s]: %s", msg, err)
		}

		return data, errorData.Content
	}

	return data, nil
}
