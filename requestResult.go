package polyanalyst6api

// RequestResult contains results of a request
type RequestResult struct {
	Body      []byte
	Additions struct {
		ExecutionWaveID *int
	}
}
