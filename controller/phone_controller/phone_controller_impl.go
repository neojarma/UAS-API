package phonecontroller

import (
	"fmt"
	"net/http"
	"uas_neoj/model/request"
	"uas_neoj/model/response"
	phoneservice "uas_neoj/service/phone_service"

	"github.com/labstack/echo/v4"
)

type PhoneControllerImpl struct {
	PhoneService phoneservice.PhoneService
}

func NewPhoneController(phoneService phoneservice.PhoneService) PhoneController {
	return &PhoneControllerImpl{
		PhoneService: phoneService,
	}
}

func (controller *PhoneControllerImpl) AllPhone(ctx echo.Context) error {
	result, err := controller.PhoneService.AllPhone()

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &response.GeneralReponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, &response.GeneralReponse{
		Code:    http.StatusOK,
		Message: "success get data",
		Data:    result,
	})
}

func (controller *PhoneControllerImpl) FindPhoneByType(ctx echo.Context) error {
	typeRequest := ctx.QueryParam("type")
	result, err := controller.PhoneService.FindPhoneByType(typeRequest)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &response.GeneralReponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, &response.GeneralReponse{
		Code:    http.StatusOK,
		Message: "success get data",
		Data:    result,
	})
}

func (controller *PhoneControllerImpl) FindPhoneBySerialNumber(ctx echo.Context) error {
	typeRequest := ctx.QueryParam("serial")

	fmt.Println(typeRequest)
	result, err := controller.PhoneService.FindPhoneBySerialNumber(typeRequest)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &response.GeneralReponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, &response.GeneralReponse{
		Code:    http.StatusOK,
		Message: "success get data",
		Data:    result,
	})
}

func (controller *PhoneControllerImpl) CreateProductionPhone(ctx echo.Context) error {
	bodyRequest := new(request.PhoneRequest)
	if err := ctx.Bind(bodyRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, &response.GeneralReponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	result, err := controller.PhoneService.CreateProductionPhone(bodyRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &response.GeneralReponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.JSON(http.StatusOK, &response.GeneralReponse{
		Code:    http.StatusOK,
		Message: "success create data",
		Data:    result,
	})

}

func (controller *PhoneControllerImpl) EditProductionPhone(ctx echo.Context) error {
	bodyRequest := new(request.PhoneRequest)
	if err := ctx.Bind(bodyRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, &response.GeneralReponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	result, err := controller.PhoneService.EditProductionPhone(bodyRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &response.GeneralReponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.JSON(http.StatusOK, &response.GeneralReponse{
		Code:    http.StatusOK,
		Message: "success update data",
		Data:    result,
	})
}

func (controller *PhoneControllerImpl) DeleteProductionPhone(ctx echo.Context) error {
	productionId := ctx.Param("productionId")
	err := controller.PhoneService.DeleteProductionPhone(productionId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &response.GeneralReponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.JSON(http.StatusOK, &response.GeneralReponse{
		Code:    http.StatusOK,
		Message: "success delete data",
		Data:    productionId,
	})
}
