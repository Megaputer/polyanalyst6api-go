package configureparameters

// LinkTermsSettins represents parameters for the Link Terms node
type LinkTermsSettins struct {
	GroupName  string `json:"Group name"`
	TermsName  string `json:"Terms name"`
	Expression string `json:"Expression"`
}

func (p LinkTermsSettins) nodeType() string {
	return "TmlLinkTerms/"
}

func (p LinkTermsSettins) settings() interface{} {
	return p
}
