package phonerepository

import (
	"database/sql"
	"uas_neoj/model/domain"
)

type PhoneRepository interface {
	AllPhone(db *sql.DB) ([]*domain.PhoneDomain, error)
	FindPhoneByType(db *sql.DB, phoneType string) ([]*domain.PhoneDomain, error)
	FindPhoneBySerialNumber(db *sql.DB, phoneSerial string) (*domain.PhoneDomain, error)
	CreateProductionPhone(db *sql.DB, phone *domain.PhoneDomain) (*domain.PhoneDomain, error)
	EditProductionPhone(db *sql.DB, phone *domain.PhoneDomain) (*domain.PhoneDomain, error)
	DeleteProductionPhone(db *sql.DB, productionId string) error
}
