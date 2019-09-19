package configureparameters

// TMLLinkTermsSettins represents parameters for Link Terms node
type TMLLinkTermsSettins struct {
	GroupName  string `json:"Group name"`
	TermsName  string `json:"Terms name"`
	Expression string `json:"Expression"`
}

func (p TMLLinkTermsSettins) nodeType() string {
	return "TmlLinkTerms/"
}

func (p TMLLinkTermsSettins) settings() interface{} {
	return p
}
