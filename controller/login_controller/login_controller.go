package controller

import (
	"github.com/labstack/echo/v4"
)

type LoginController interface {
	Login(ctx echo.Context) error
}
