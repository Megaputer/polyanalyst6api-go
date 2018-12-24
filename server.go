package polyanalyst6api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	apiRoot = "/polyanalyst/api"
)

// Server is the struct for server data representation
type Server struct {
	Host string
	Port uint
}

// BaseURL returns base API url for the server
func (s Server) BaseURL() string {
	return fmt.Sprintf("https://%s:%d%s", s.Host, s.Port, apiRoot)
}

// APIVersions returns available API versions
func (s Server) APIVersions() ([]string, error) {
	var (
		vs  []string
		err error
	)

	resp, err := http.Get(s.BaseURL() + "/versions")
	defer resp.Body.Close()
	if err != nil {
		return vs, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return vs, err
	}

	err = json.Unmarshal(data, &vs)

	return vs, err
}
