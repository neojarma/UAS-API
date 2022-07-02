package loginservice

import (
	"uas_neoj/model/request"
)

type LoginService interface {
	Login(request *request.LoginRequest) error
	CreateLogin(request *request.LoginRequest) error
}
