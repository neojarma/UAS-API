package router

import (
	"database/sql"
	"uas_neoj/authentication"
	employeecontroller "uas_neoj/controller/employee_controller"
	logincontroller "uas_neoj/controller/login_controller"
	phonecontroller "uas_neoj/controller/phone_controller"
	tokencontroller "uas_neoj/controller/token_controller"
	employeerepository "uas_neoj/repository/employee_repository"
	loginrepository "uas_neoj/repository/login_repository"
	phonerepository "uas_neoj/repository/phone_repository"
	employeeservice "uas_neoj/service/employee_service"
	loginservice "uas_neoj/service/login_service"
	phoneservice "uas_neoj/service/phone_service"

	"github.com/labstack/echo/v4"
)

func Router(group *echo.Group, db *sql.DB) {
	loginRouter(group, db)
	employeeRouter(group, db)
	phoneRouter(group, db)
	docs(group)
	token(group)
}

func token(group *echo.Group) {
	tokenController := tokencontroller.NewTokenController()
	group.POST("/token", tokenController.RefreshToken)
}

func phoneRouter(group *echo.Group, db *sql.DB) {
	phoneRepo := phonerepository.NewPhoneRepository()
	phoneService := phoneservice.NewPhoneService(phoneRepo, db)
	phoneController := phonecontroller.NewPhoneController(phoneService)

	group.GET("/phones", phoneController.AllPhone)
	group.GET("/phones/type", phoneController.FindPhoneByType)
	group.GET("/phone/serial", phoneController.FindPhoneBySerialNumber)
	group.POST("/phone", phoneController.CreateProductionPhone, authentication.IsLoggedIn)
	group.PUT("/phone", phoneController.EditProductionPhone, authentication.IsLoggedIn)
	group.DELETE("/phone/:productionId", phoneController.DeleteProductionPhone, authentication.IsLoggedIn)
}

func loginRouter(group *echo.Group, db *sql.DB) {
	loginRepo := loginrepository.NewLoginRepository()
	loginService := loginservice.NewLoginService(loginRepo, db)
	loginController := logincontroller.NewLoginController(loginService)

	group.POST("/login", loginController.Login)
}

func employeeRouter(group *echo.Group, db *sql.DB) {
	employeeRepo := employeerepository.NewEmployeeRepository()
	loginRepo := loginrepository.NewLoginRepository()
	employeeService := employeeservice.NewEmployeeService(employeeRepo, loginRepo, db)
	employeeController := employeecontroller.NewEmployeeController(employeeService)

	group.POST("/employee/register", employeeController.Register)
}

func docs(group *echo.Group) {
	group.Static("/docs", "./docs")
}
