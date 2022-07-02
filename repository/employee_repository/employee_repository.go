package employeerepository

import (
	"database/sql"
	"uas_neoj/model/domain"
)

type EmployeeRepository interface {
	Register(db *sql.DB, employee *domain.EmployeeDomain) (*domain.EmployeeDomain, error)
}
