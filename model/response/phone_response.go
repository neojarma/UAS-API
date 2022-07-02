package response

type PhoneResponse struct {
	IdProduction   string `json:"id_production"`
	PhoneName      string `json:"phone_name"`
	PhoneType      string `json:"phone_type"`
	SerialNumber   string `json:"serial_number"`
	ProductionDate string `json:"production_date"`
}
