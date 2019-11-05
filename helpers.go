package polyanalyst6api

import (
	"net/http"
)

func closeBody(r *http.Response) {
	_ = r.Body.Close()
}

func isMain() bool {
	return branch == "main"
}

func boolToURLParam(b bool) string {
	if b {
		return "1"
	}
	return "0"
}
