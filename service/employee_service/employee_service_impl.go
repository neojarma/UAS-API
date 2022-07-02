package employeeservice

import (
	"database/sql"
	"uas_neoj/helper"
	"uas_neoj/model/domain"
	"uas_neoj/model/request"
	employeerepository "uas_neoj/repository/employee_repository"
	loginrepository "uas_neoj/repository/login_repository"
)

type EmployeeServiceImpl struct {
	EmployeeRepository employeerepository.EmployeeRepository
	LoginRepository    loginrepository.LoginRepository
	Db                 *sql.DB
}

func NewEmployeeService(repo employeerepository.EmployeeRepository, login loginrepository.LoginRepository, db *sql.DB) EmployeeService {
	return &EmployeeServiceImpl{
		EmployeeRepository: repo,
		LoginRepository:    login,
		Db:                 db,
	}
}

func (service *EmployeeServiceImpl) Register(request *request.EmployeeRequest) error {
	idEmployee := helper.GenerateRandomId("USR-")

	employee := domain.EmployeeDomain{
		IdEmployee:  idEmployee,
		FullName:    request.FullName,
		Address:     request.Address,
		PhoneNumber: request.PhoneNumber,
		Username:    request.Username,
	}

	_, err := service.EmployeeRepository.Register(service.Db, &employee)

	if err != nil {
		return err
	}

	login := domain.LoginDomain{
		Username: request.Username,
		Password: request.Password,
	}

	_, err = service.LoginRepository.CreateLogin(service.Db, &login)
	if err != nil {
		return err
	}

	return nil
}
