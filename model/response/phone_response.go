package response

type PhoneResponse struct {
	IdProduction   string `json:"idProduction"`
	PhoneName      string `json:"phoneName"`
	PhoneType      string `json:"phoneType"`
	SerialNumber   string `json:"serialNumber"`
	ProductionDate string `json:"productionDate"`
}
