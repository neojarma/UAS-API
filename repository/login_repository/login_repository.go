package loginrepository

import (
	"database/sql"
	"uas_neoj/model/domain"
)

type LoginRepository interface {
	Login(db *sql.DB, login *domain.LoginDomain) (*domain.LoginDomain, error)
	CreateLogin(db *sql.DB, login *domain.LoginDomain) (*domain.LoginDomain, error)
}
