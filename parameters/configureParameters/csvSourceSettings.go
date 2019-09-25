package configureparameters

// CSVSourceSettings represents parameters for the CSV Source node
type CSVSourceSettings struct {
	File string `json:"File"`
}

func (s CSVSourceSettings) nodeType() string {
	return "DataSource/CSV"
}

func (s CSVSourceSettings) settings() interface{} {
	return s
}
