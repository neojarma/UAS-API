package employeecontroller

import (
	"github.com/labstack/echo/v4"
)

type EmployeeController interface {
	Register(ctx echo.Context) error
}
