package polyanalyst6api

// SetAlphaErrorProcessing sets the alpha mode (new error structure)
func SetAlphaErrorProcessing() {
	branch = "alpha"
}

// SetMainErrorProcessing sets the main mode
func SetMainErrorProcessing() {
	branch = "main"
}

// IsPABUSY determines if this is a PABUSY error
func IsPABUSY(err error) bool {
	serverError, ok := err.(ServerError)
	if !ok {
		return false
	}
	return serverError.Code == 503
}
