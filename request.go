package polyanalyst6api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
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

func (r request) Perform() (RequestResult, error) {
	var result RequestResult

	client := &http.Client{
		Timeout: RequestTimeout,
	}

	resp, err := client.Do(r.httpReq)
	if err != nil {
		return result, fmt.Errorf("request execution error: %w", err)
	}
	defer closeBody(resp)

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		var errorData serverErrorData
		err = json.Unmarshal(data, &errorData)
		if err != nil {
			return result, fmt.Errorf("failed to parse server error [%s]: %w", data, err)
		}

		return result, errorData.Content
	}

	result.Body = data
	locURL, err := resp.Location()
	if err == nil {
		params, err := url.ParseQuery(locURL.RawQuery)
		if err == nil {
			execWaveStrings, ok := params["executionWave"]
			if ok && len(execWaveStrings) > 0 {
				execWaveInt, err := strconv.Atoi(execWaveStrings[0])
				if err == nil {
					result.Additions.ExecutionWaveID = new(int)
					result.Additions.ExecutionWaveID = &execWaveInt
				}
			}
		}
	}

	return result, nil
}
