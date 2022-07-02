package employeeservice

import (
	"uas_neoj/model/request"
)

type EmployeeService interface {
	Register(request *request.EmployeeRequest) error
}
