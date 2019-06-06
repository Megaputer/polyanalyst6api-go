package polyanalyst6api

import (
	"bytes"
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

	data, errBodyRead := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		msg := ""
		if errBodyRead != nil {
			msg = "*failed to retrieve*"
		} else {
			msg = string(data)
		}

		fullMsg := fmt.Sprintf("bad response status: %d. Error: %s", resp.StatusCode, msg)

		if resp.StatusCode == 503 {
			return data, PABUSY{err: fullMsg}
		}

		return data, fmt.Errorf("bad response status: %d. Error: %s", resp.StatusCode, msg)
	}

	if errBodyRead != nil {
		return data, fmt.Errorf("failed to read response")
	}

	return data, nil
}
