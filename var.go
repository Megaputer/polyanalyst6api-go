package polyanalyst6api

import "time"

var (
	// RequestTimeout defines the request timeout
	// It is passed to &http.Client{Timeout: RequestTimeout}
	// IMPORTANT: don't set it to 0 (it may lead to endless waitings)
	RequestTimeout = 1 * time.Second

	branch = "main"
)
