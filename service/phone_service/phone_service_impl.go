package phoneservice

import (
	"database/sql"
	"uas_neoj/helper"
	"uas_neoj/model/domain"
	"uas_neoj/model/request"
	"uas_neoj/model/response"
	phonerepository "uas_neoj/repository/phone_repository"
)

type phoneServiceImpl struct {
	PhoneRepository phonerepository.PhoneRepository
	Db              *sql.DB
}

func NewPhoneService(repo phonerepository.PhoneRepository, db *sql.DB) PhoneService {
	return &phoneServiceImpl{
		PhoneRepository: repo,
		Db:              db,
	}
}

func (service *phoneServiceImpl) AllPhone() ([]*response.PhoneResponse, error) {
	domainResult, err := service.PhoneRepository.AllPhone(service.Db)

	if err != nil {
		return nil, err
	}

	var result []*response.PhoneResponse
	for _, domain := range domainResult {
		result = append(result, &response.PhoneResponse{
			IdProduction:   domain.IdProduction,
			PhoneName:      domain.PhoneName,
			PhoneType:      domain.PhoneType,
			SerialNumber:   domain.SerialNumber,
			ProductionDate: domain.ProductionDate,
		})
	}

	return result, nil
}

func (service *phoneServiceImpl) FindPhoneByType(phoneType string) ([]*response.PhoneResponse, error) {
	domainResult, err := service.PhoneRepository.FindPhoneByType(service.Db, phoneType)

	if err != nil {
		return nil, err
	}

	var result []*response.PhoneResponse
	for _, domain := range domainResult {
		result = append(result, &response.PhoneResponse{
			IdProduction:   domain.IdProduction,
			PhoneName:      domain.PhoneName,
			PhoneType:      domain.PhoneType,
			SerialNumber:   domain.SerialNumber,
			ProductionDate: domain.ProductionDate,
		})
	}

	return result, nil
}

func (service *phoneServiceImpl) FindPhoneBySerialNumber(phoneSerial string) (*response.PhoneResponse, error) {
	domainResult, err := service.PhoneRepository.FindPhoneBySerialNumber(service.Db, phoneSerial)

	if err != nil {
		return nil, err
	}

	result := &response.PhoneResponse{
		IdProduction:   domainResult.IdProduction,
		PhoneName:      domainResult.PhoneName,
		PhoneType:      domainResult.PhoneType,
		SerialNumber:   domainResult.SerialNumber,
		ProductionDate: domainResult.ProductionDate,
	}

	return result, nil
}

func (service *phoneServiceImpl) CreateProductionPhone(request *request.PhoneRequest) (*response.PhoneResponse, error) {

	idProduction := helper.GenerateRandomId("PROD-")

	domainRequest := &domain.PhoneDomain{
		IdProduction:   idProduction,
		PhoneName:      request.PhoneName,
		PhoneType:      request.PhoneType,
		SerialNumber:   request.SerialNumber,
		ProductionDate: request.ProductionDate,
	}

	domainResult, err := service.PhoneRepository.CreateProductionPhone(service.Db, domainRequest)

	if err != nil {
		return nil, err
	}

	result := &response.PhoneResponse{
		IdProduction:   domainResult.IdProduction,
		PhoneName:      domainResult.PhoneName,
		PhoneType:      domainResult.PhoneType,
		SerialNumber:   domainResult.SerialNumber,
		ProductionDate: domainResult.ProductionDate,
	}

	return result, nil
}

func (service *phoneServiceImpl) EditProductionPhone(request *request.PhoneRequest) (*response.PhoneResponse, error) {

	domainRequest := &domain.PhoneDomain{
		IdProduction:   request.IdProduction,
		PhoneName:      request.PhoneName,
		PhoneType:      request.PhoneType,
		SerialNumber:   request.SerialNumber,
		ProductionDate: request.ProductionDate,
	}

	domainResult, err := service.PhoneRepository.EditProductionPhone(service.Db, domainRequest)

	if err != nil {
		return nil, err
	}

	result := &response.PhoneResponse{
		IdProduction:   domainResult.IdProduction,
		PhoneName:      domainResult.PhoneName,
		PhoneType:      domainResult.PhoneType,
		SerialNumber:   domainResult.SerialNumber,
		ProductionDate: domainResult.ProductionDate,
	}

	return result, nil

}

func (service *phoneServiceImpl) DeleteProductionPhone(productionId string) error {
	err := service.PhoneRepository.DeleteProductionPhone(service.Db, productionId)

	if err != nil {
		return err
	}

	return nil

}
