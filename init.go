package polyanalyst6api

import (
	"crypto/tls"
	"net/http"
)

func init() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
}
