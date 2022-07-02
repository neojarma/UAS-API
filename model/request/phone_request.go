package request

type PhoneRequest struct {
	IdProduction   string `json:"idProduction,omitempty"`
	PhoneName      string `json:"phoneName"`
	PhoneType      string `json:"phoneType"`
	SerialNumber   string `json:"serialNumber"`
	ProductionDate string `json:"productionDate"`
}
