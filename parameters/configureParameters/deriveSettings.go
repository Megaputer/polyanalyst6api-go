package configureparameters

// DeriveSettings represents parameters for the Derive node
type DeriveSettings struct {
	Name string `json:"Name"`
	Rule string `json:"Rule"`
}

func (s DeriveSettings) nodeType() string {
	return "SRLRuleSet/SRL Rule"
}

func (s DeriveSettings) settings() interface{} {
	return s
}
