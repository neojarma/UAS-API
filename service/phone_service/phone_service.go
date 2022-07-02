package phoneservice

import (
	"uas_neoj/model/request"
	"uas_neoj/model/response"
)

type PhoneService interface {
	AllPhone() ([]*response.PhoneResponse, error)
	FindPhoneByType(phoneType string) ([]*response.PhoneResponse, error)
	FindPhoneBySerialNumber(phoneSerial string) (*response.PhoneResponse, error)
	CreateProductionPhone(request *request.PhoneRequest) (*response.PhoneResponse, error)
	EditProductionPhone(request *request.PhoneRequest) (*response.PhoneResponse, error)
	DeleteProductionPhone(productionId string) error
}
