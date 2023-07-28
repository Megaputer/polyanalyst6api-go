package responses

type ServerHealth struct {
	Status  string `json:"status"`
	Details struct {
		License struct {
			Status string `json:"status"`
		} `json:"license"`
		MaintenanceMode struct {
			Status string `json:"status"`
		} `json:"maintenanceMode"`
	} `json:"details"`
}
