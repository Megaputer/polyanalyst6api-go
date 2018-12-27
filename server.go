package polyanalyst6api

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gluk-skywalker/polyanalyst6api-go/parameters"
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

	url := s.BaseURL() + "/versions"
	req, err := CreateRequest(url, "GET", parameters.Full{})
	if err != nil {
		return vs, err
	}

	data, err := req.Perform()
	if err != nil {
		return vs, err
	}

	err = json.Unmarshal(data, &vs)

	return vs, err
}

// SupportsAPIVersion checks if tha passed API version is supported by the server
func (s Server) SupportsAPIVersion(v string) (bool, error) {
	var (
		res bool
		err error
	)

	supportedVersions, err := s.APIVersions()
	if err != nil {
		return res, errors.New("failed to get supported API versions: " + err.Error())
	}

	for _, srvVersion := range supportedVersions {
		if srvVersion == v {
			return true, nil
		}
	}
	return false, nil
}
