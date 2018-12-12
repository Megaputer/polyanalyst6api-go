package parameters

import (
	"net/url"
)

// FullParams is struct that is suitable for any type of a request: GET/POST
type FullParams struct {
	URLParams  url.Values
	BodyParams []byte
}
