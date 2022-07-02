package employeecontroller

import (
	"net/http"
	"uas_neoj/model/request"
	"uas_neoj/model/response"
	employeeservice "uas_neoj/service/employee_service"

	"github.com/labstack/echo/v4"
)

type EmployeeControllerImpl struct {
	EmployeeService employeeservice.EmployeeService
}

func NewEmployeeController(employeeService employeeservice.EmployeeService) EmployeeController {
	return &EmployeeControllerImpl{
		EmployeeService: employeeService,
	}
}

func (controller *EmployeeControllerImpl) Register(ctx echo.Context) error {
	bodyRequest := new(request.EmployeeRequest)
	if err := ctx.Bind(bodyRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, &response.GeneralReponse{
			Code:    http.StatusBadRequest,
			Message: "Bad Request",
			Data:    nil,
		})
	}

	if err := controller.EmployeeService.Register(bodyRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, &response.GeneralReponse{
			Code:    http.StatusBadRequest,
			Message: "Bad Request",
			Data:    nil,
		})
	}

	return ctx.JSON(http.StatusOK, &response.GeneralReponse{
		Code:    http.StatusOK,
		Message: "success register",
		Data:    bodyRequest,
	})
}
