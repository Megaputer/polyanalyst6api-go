package objects

// Represents the structure for server info
type ServerInfo struct {
	Build   int    `json:"build"`
	Version string `json:"version"`
	GitSHA1 struct {
		Dicts string `json:"dicts"`
		Help  string `json:"help"`
		PA    string `json:"pa"`
		Webui string `json:"webui"`
	} `json:"gitSHA1"`
}
