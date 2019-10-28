package polyanalyst6api

import "fmt"

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
	return fmt.Sprintf("%d: %s: %s", e.Code, e.Title, e.Msg())
}
