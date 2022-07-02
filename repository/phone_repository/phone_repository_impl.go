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

func (repository *phoneRepositoryImpl) AllPhone(db *sql.DB) ([]*domain.PhoneDomain, error) {
	SQL := "SELECT id_production, phone_name, phone_type, serial_number, production_date FROM public.phones"
	rows, err := db.Query(SQL)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var phones []*domain.PhoneDomain

	for rows.Next() {
		var phone domain.PhoneDomain
		err := rows.Scan(&phone.IdProduction, &phone.PhoneName, &phone.PhoneType, &phone.SerialNumber, &phone.ProductionDate)
		if err != nil {
			return nil, err
		}
		phones = append(phones, &phone)
	}

	return phones, nil
}

func (repository *phoneRepositoryImpl) FindPhoneByType(db *sql.DB, phoneType string) ([]*domain.PhoneDomain, error) {
	SQL := "SELECT id_production, phone_name, phone_type, serial_number, production_date FROM public.phones WHERE phone_type = $1"
	rows, err := db.Query(SQL, phoneType)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var phones []*domain.PhoneDomain

	for rows.Next() {
		var phone domain.PhoneDomain
		err := rows.Scan(&phone.IdProduction, &phone.PhoneName, &phone.PhoneType, &phone.SerialNumber, &phone.ProductionDate)
		if err != nil {
			return nil, err
		}
		phones = append(phones, &phone)
	}

	return phones, nil
}

func (repository *phoneRepositoryImpl) FindPhoneBySerialNumber(db *sql.DB, phoneSerial string) (*domain.PhoneDomain, error) {
	SQL := "SELECT id_production, phone_name, phone_type, serial_number, production_date FROM public.phones WHERE serial_number = $1"
	rows, err := db.Query(SQL, phoneSerial)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var phone domain.PhoneDomain
	if rows.Next() {
		err := rows.Scan(&phone.IdProduction, &phone.PhoneName, &phone.PhoneType, &phone.SerialNumber, &phone.ProductionDate)
		if err != nil {
			return nil, err
		}
	}

	return &phone, nil
}

func (repository *phoneRepositoryImpl) CreateProductionPhone(db *sql.DB, phone *domain.PhoneDomain) (*domain.PhoneDomain, error) {
	SQL := "INSERT INTO public.phones (id_production, phone_name, phone_type, serial_number, production_date) VALUES($1, $2, $3, $4, $5)"
	_, err := db.Exec(SQL, phone.IdProduction, phone.PhoneName, phone.PhoneType, phone.SerialNumber, phone.ProductionDate)
	if err != nil {
		return nil, err
	}

	return phone, nil
}

func (repository *phoneRepositoryImpl) EditProductionPhone(db *sql.DB, phone *domain.PhoneDomain) (*domain.PhoneDomain, error) {
	SQL := "UPDATE public.phones SET phone_name = $1, phone_type = $2, serial_number = $3, production_date = $4 WHERE id_production = $5"
	_, err := db.Exec(SQL, phone.PhoneName, phone.PhoneType, phone.SerialNumber, phone.ProductionDate, phone.IdProduction)
	if err != nil {
		return nil, err
	}

	return phone, nil
}

func (repository *phoneRepositoryImpl) DeleteProductionPhone(db *sql.DB, productionId string) error {
	SQL := "DELETE FROM public.phones WHERE id_production = $1"
	_, err := db.Exec(SQL, productionId)
	if err != nil {
		return nil
	}

	return nil
}
