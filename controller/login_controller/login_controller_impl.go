package controller

import (
	"net/http"
	"uas_neoj/model/request"
	"uas_neoj/model/response"
	loginservice "uas_neoj/service/login_service"

	"github.com/labstack/echo/v4"
)

type LoginControllerImpl struct {
	LoginService loginservice.LoginService
}

func NewLoginController(service loginservice.LoginService) LoginController {
	return &LoginControllerImpl{
		LoginService: service,
	}
}

func (controller *LoginControllerImpl) Login(ctx echo.Context) error {
	bodyRequest := new(request.LoginRequest)

	if err := ctx.Bind(bodyRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, &response.GeneralReponse{
			Code:    http.StatusBadRequest,
			Message: "Bad Request",
			Data:    nil,
		})
	}

	if err := controller.LoginService.Login(bodyRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, &response.GeneralReponse{
			Code:    http.StatusBadRequest,
			Message: "Bad Request",
			Data:    nil,
		})
	}

	return ctx.JSON(http.StatusOK, &response.GeneralReponse{
		Code:    http.StatusOK,
		Message: "access granted",
		Data:    nil,
	})
}
