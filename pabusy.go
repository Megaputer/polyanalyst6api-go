package polyanalyst6api

// PABUSY is a special kind of error (polyanalyst server is busy)
type PABUSY struct {
	err string
}

func (e PABUSY) Error() string {
	return e.err
}
