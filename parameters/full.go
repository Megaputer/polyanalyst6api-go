package parameters

import (
	"net/url"
)

// FullParams is struct that is suitable for any type of a request: GET/POST
type Full struct {
	URLParams  url.Values
	BodyParams []byte
}

// Мне не очень понятна идея с параметрами. Можешь объяснить?
