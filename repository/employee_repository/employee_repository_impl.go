package employeerepository

import (
	"database/sql"
	"uas_neoj/model/domain"
)

type employeeRepositoryImpl struct {
}

func NewEmployeeRepository() EmployeeRepository {
	return &employeeRepositoryImpl{}
}

func (repository *employeeRepositoryImpl) Register(db *sql.DB, employee *domain.EmployeeDomain) (*domain.EmployeeDomain, error) {
	SQL := "INSERT INTO public.employee (id_employee, full_name, address, phone_number, username) VALUES ($1, $2, $3, $4, $5)"
	_, err := db.Exec(SQL, employee.IdEmployee, employee.FullName, employee.Address, employee.PhoneNumber, employee.Username)
	if err != nil {
		return nil, err
	}

	return employee, nil
}
