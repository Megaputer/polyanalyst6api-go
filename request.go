package polyanalyst6api

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gluk-skywalker/polyanalyst6api-go/parameters"
)

type request struct {
	session *Session
	path    string
	reqType string
	params  parameters.Full
}

func (r request) Perform() ([]byte, error) {
	var (
		err  error
		data []byte
	)

	// turning url paras to RFC 3986 compatible string
	urlParams := strings.Replace(r.params.URLParams.Encode(), "+", "%20", -1)

	url := r.path + "?" + urlParams
	req, err := http.NewRequest(r.reqType, url, bytes.NewBuffer(r.params.BodyParams))
	if err != nil {
		return data, errors.New("building request error: " + err.Error())
	}

	cookie := http.Cookie{Name: "sid", Value: r.session.SID}
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
