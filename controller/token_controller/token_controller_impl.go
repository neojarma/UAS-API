package tokencontroller

import (
	"errors"
	"net/http"
	"uas_neoj/authentication"
	"uas_neoj/helper"
	"uas_neoj/model/request"
	"uas_neoj/model/response"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type tokenControllerImpl struct {
}

func NewTokenController() TokenController {
	return &tokenControllerImpl{}
}

func (controller *tokenControllerImpl) RefreshToken(ctx echo.Context) error {
	request := new(request.RefreshTokenRequest)
	if err := ctx.Bind(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, &response.GeneralReponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	userRefreshToken := request.RefreshToken
	token, err := jwt.Parse(userRefreshToken, func(t *jwt.Token) (interface{}, error) {
		return []byte(helper.LoadEnv("REFRESH_TOKEN_SECRET")), nil
	})

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &response.GeneralReponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	if !token.Valid {
		return ctx.JSON(http.StatusBadRequest, &response.GeneralReponse{
			Code:    http.StatusBadRequest,
			Message: errors.New("invalid refresh token").Error(),
		})
	}

	newAccessToken, err := authentication.GenerateTokenPair()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &response.GeneralReponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, &response.GeneralReponse{
		Code:    http.StatusOK,
		Message: "success refresh access token",
		Data:    newAccessToken,
	})
}
