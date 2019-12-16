package polyanalyst6api

import (
	"strconv"
)

// ServerError represents PolyAnalyst server error
type ServerError struct {
	Code    int    `json:"code"`
	Title   string `json:"title"`
	Message string `json:"message"`
}

// Msg returns message if there is and [empty] otherwise
func (e ServerError) Msg() string {
	if len(e.Message) == 0 {
		return "[empty]"
	}
	return e.Message
}

func (e ServerError) Error() string {
	msg := strconv.Itoa(e.Code)
	if len(e.Title) > 0 {
		msg += ": " + e.Title
	}
	if len(e.Message) > 0 {
		msg += ": " + e.Message
	}
	return msg
}
