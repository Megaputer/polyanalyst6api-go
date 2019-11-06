package responses

type ParametersAvailableNode struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	Parameters []struct {
		Name     string `json:"name"`
		Type     string `json:"type"`
		Optional int    `json:"optional"`
	} `json:"parameters"`
}
