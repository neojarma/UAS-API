package domain

type EmployeeDomain struct {
	IdEmployee  string `json:"id_employee"`
	FullName    string `json:"full_name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone"`
	Username    string `json:"username"`
}
