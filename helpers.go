package polyanalyst6api

import (
	"net/http"
	"os"
)

func pathExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func closeBody(r *http.Response) {
	_ = r.Body.Close()
}
