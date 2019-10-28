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
