package phonecontroller

import (
	"github.com/labstack/echo/v4"
)

type PhoneController interface {
	AllPhone(ctx echo.Context) error
	FindPhoneByType(ctx echo.Context) error
	FindPhoneBySerialNumber(ctx echo.Context) error
	CreateProductionPhone(ctx echo.Context) error
	EditProductionPhone(ctx echo.Context) error
	DeleteProductionPhone(ctx echo.Context) error
}
