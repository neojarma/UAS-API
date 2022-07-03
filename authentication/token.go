package authentication

import (
	"time"
	"uas_neoj/helper"

	"github.com/golang-jwt/jwt"
)

func GenerateTokenPair() (map[string]string, error) {
	accessSecret := helper.LoadEnv("ACCESS_TOKEN_SECRET")
	refreshSecret := helper.LoadEnv("REFRESH_TOKEN_SECRET")

	accessClaims := jwt.StandardClaims{
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
	}
	tokenAccess := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)

	accessToken, err := tokenAccess.SignedString([]byte(accessSecret))
	if err != nil {
		return nil, err
	}

	refreshClaims := jwt.StandardClaims{
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}
	tokenRefresh := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshToken, err := tokenRefresh.SignedString([]byte(refreshSecret))
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"accessIssuedAt":   time.Unix(accessClaims.IssuedAt, 0).String(),
		"accessExpiresAt":  time.Unix(accessClaims.ExpiresAt, 0).String(),
		"refreshExpiresAt": time.Unix(refreshClaims.ExpiresAt, 0).String(),
		"accessToken":      accessToken,
		"refreshToken":     refreshToken,
	}, nil
}
