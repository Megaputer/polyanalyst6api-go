package polyanalyst6api

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// InitSession is used to start a session
func InitSession(server *Server, version string, login string, password string) (Session, error) {
	var session Session

	supports, err := server.SupportsAPIVersion(version)
	if err != nil {
		return session, errors.New("failed to check if the passed API version is supported: " + err.Error())
	}

	if !supports {
		return session, errors.New("the server doesn't support API of version " + version)
	}

	url := server.BaseURL() + fmt.Sprintf("/v%s/login?uname=%s&pwd=%s", version, login, password)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return session, errors.New("building request error: " + err.Error())
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return session, errors.New("request execution error: " + err.Error())
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
		return session, fmt.Errorf("bad response status: %d. Error: %s", resp.StatusCode, msg)
	}

	for _, c := range resp.Cookies() {
		if c.Name == "sid" {
			return Session{SID: c.Value, Server: server, apiVersion: version}, nil
		}
	}

	return session, errors.New("login response does not contain the sid")
}
