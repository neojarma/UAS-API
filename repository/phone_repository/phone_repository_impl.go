package phonerepository

import (
	"database/sql"
	"uas_neoj/model/domain"
)

type phoneRepositoryImpl struct {
}

func NewPhoneRepository() PhoneRepository {
	return &phoneRepositoryImpl{}
}

func (repository *phoneRepositoryImpl) AllPhone(db *sql.DB, phone *domain.PhoneDomain) ([]*domain.PhoneDomain, error) {
	panic("not implemented") // TODO: Implement
}

func (repository *phoneRepositoryImpl) FindPhoneByType(db *sql.DB, phoneType string) (*domain.PhoneDomain, error) {
	panic("not implemented") // TODO: Implement
}

func (repository *phoneRepositoryImpl) FindPhoneBySerialNumber(db *sql.DB, phoneSerial string) (*domain.PhoneDomain, error) {
	panic("not implemented") // TODO: Implement
}

func (repository *phoneRepositoryImpl) CreateProductionPhone(db *sql.DB, phone *domain.PhoneDomain) error {
	panic("not implemented") // TODO: Implement
}

func (repository *phoneRepositoryImpl) EditProductionPhone(db *sql.DB, phone *domain.PhoneDomain) error {
	panic("not implemented") // TODO: Implement
}

func (repository *phoneRepositoryImpl) DeleteProductionPhone(db *sql.DB, phone *domain.PhoneDomain) error {
	panic("not implemented") // TODO: Implement
}
