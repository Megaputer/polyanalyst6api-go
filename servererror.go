package polyanalyst6api

import (
	"strconv"
)

// ServerError represents PolyAnalyst server error
type ServerError struct {
	Info struct {
		Code    int    `json:"code"`
		Title   string `json:"title"`
		Message string `json:"message"`
	} `json:"error"`
}

// Msg returns message if there is and [empty] otherwise
func (e ServerError) Msg() string {
	if len(e.Info.Message) == 0 {
		return "[empty]"
	}
	return e.Info.Message
}

func (e ServerError) Error() string {
	msg := strconv.Itoa(e.Info.Code)
	if len(e.Info.Title) > 0 {
		msg += ": " + e.Info.Title
	}
	if len(e.Info.Message) > 0 {
		msg += ": " + e.Info.Message
	}
	return msg
}
