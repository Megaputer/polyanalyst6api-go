package polyanalyst6api

// SetAlphaErrorProcessing sets the alpha mode (new error structure)
func SetAlphaErrorProcessing() {
	branch = "alpha"
}

// SetMainErrorProcessing sets the main mode
func SetMainErrorProcessing() {
	branch = "main"
}

// IsPABUSYError determines if this is a PABUSY error
func IsPABUSYError(err error) bool {
	serverError, ok := err.(ServerError)
	if !ok {
		return false
	}
	return serverError.Code == 503
}

// IsDBRecoveryError determines if this is a DBRecovery error
func IsDBRecoveryError(err error) bool {
	serverError, ok := err.(ServerError)
	if !ok {
		return false
	}
	msg := "PolyAnalyst Server performs database recovery and cannot establish connection right now. Please wait a minute and retry operation"
	return serverError.Message == msg
}
