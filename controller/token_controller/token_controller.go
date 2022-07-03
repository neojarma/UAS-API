package tokencontroller

import "github.com/labstack/echo/v4"

type TokenController interface {
	RefreshToken(ctx echo.Context) error
}
