package polyanalyst6api

import (
	"encoding/json"
	"fmt"
	"io"

	"net/http"
)

// InitSession is used to start a session
func InitSession(server *Server, version string, login string, password string, useLDAP bool, LDAPServer string) (Session, error) {
	var session Session

	supports, err := server.SupportsAPIVersion(version)
	if err != nil {
		return session, fmt.Errorf("failed to check if the passed API version is supported: %w", err)
	}

	if !supports {
		return session, fmt.Errorf("the server doesn't support API of version %s", version)
	}

	url := server.BaseURL() + fmt.Sprintf("/v%s/login?uname=%s&pwd=%s&useLDAP=%s&svr=%s", version, login, password, boolToURLParam(useLDAP), LDAPServer)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return session, fmt.Errorf("building request error: %w", err)
	}

	client := &http.Client{
		Timeout: RequestTimeout,
	}

	resp, err := client.Do(req)
	if err != nil {
		return session, fmt.Errorf("request execution error: %w", err)
	}
	defer closeBody(resp)

	if resp.StatusCode != 200 {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return session, fmt.Errorf("failed to read request response: %w", err)
		}

		var serverErr ServerError
		err = json.Unmarshal(bodyBytes, &serverErr)
		if err != nil {
			return session, fmt.Errorf("failed to parse request response: %w", err)
		}

		return session, fmt.Errorf("bad response status: %d. Error: %w", resp.StatusCode, serverErr)
	}

	for _, c := range resp.Cookies() {
		if c.Name == "sid" {
			return Session{SID: c.Value, Server: server, apiVersion: version}, nil
		}
	}

	return session, fmt.Errorf("login response does not contain the sid")
}
