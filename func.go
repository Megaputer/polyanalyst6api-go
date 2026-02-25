package polyanalyst6api

import "errors"

// IsPABUSYError determines if this is a PABUSY error
func IsPABUSYError(err error) bool {
	var serverErr ServerError
	ok := errors.As(err, &serverErr)
	if !ok {
		return false
	}

	return serverErr.Info.Code == 503
}

// IsDBRecoveryError determines if this is a DBRecovery error
func IsDBRecoveryError(err error) bool {
	var serverErr ServerError
	ok := errors.As(err, &serverErr)
	if !ok {
		return false
	}

	msg := "PolyAnalyst Server performs database recovery and cannot establish connection right now. Please wait a minute and retry operation"
	return serverErr.Info.Message == msg
}
