package polyanalyst6api

// PABUSY is a special kind of error (polyanalyst server is busy)
type PABUSY struct {
	err string
}

func (e PABUSY) Error() string {
	return e.err
}

// IsPABUSY determines if this is a PABUSY error
func IsPABUSY(err error) bool {
	_, ok := err.(PABUSY)
	return ok
}
