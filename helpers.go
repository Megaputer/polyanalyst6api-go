package polyanalyst6api

import (
	"net/http"
)

func closeBody(r *http.Response) {
	_ = r.Body.Close()
}

func boolToURLParam(b bool) string {
	if b {
		return "1"
	}
	return "0"
}
