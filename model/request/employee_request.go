package request

type EmployeeRequest struct {
	FullName    string `json:"fullName"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone"`
	Username    string `json:"username"`
	Password    string `json:"password"`
}
