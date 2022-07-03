package authentication

import (
	"uas_neoj/helper"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4/middleware"
)

var IsLoggedIn = middleware.JWTWithConfig(

	middleware.JWTConfig{
		SigningMethod: jwt.SigningMethodHS256.Name,
		SigningKey:    []byte(helper.LoadEnv("ACCESS_TOKEN_SECRET")),
	},
)
