package polyanalyst6api

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Session is used to interact with the API
type Session struct {
	SID     string
	BaseURL string
}

// Request is used for making requests to the API
func (s Session) Request(reqType string, path string, params string) error {
	var err error

	url := s.BaseURL + path + "?" + params
	req, err := http.NewRequest(reqType, url, nil)
	if err != nil {
		return errors.New("building request error: " + err.Error())
	}

	cookie := http.Cookie{Name: "sid", Value: s.SID}
	req.AddCookie(&cookie)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return errors.New("request execution error: " + err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		msg := ""
		if err != nil {
			msg = "*failed to retrieve*"
		} else {
			msg = string(bodyBytes)
		}
		return fmt.Errorf("bad response status: %d. Error: %s", resp.StatusCode, msg)
	}

	return nil
}
