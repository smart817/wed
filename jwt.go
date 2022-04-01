package wed

import (
	"errors"

	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte(config.Jwt.Secret) //config.AllConfig.Jwt.Secret

type myClaims struct {
	JwtId string `json:"jwtId"`
	jwt.StandardClaims
}

func CreateToken(jwtId string) string {
	if jwtId == "" {
		return "jwtId null"
	}
	claims := myClaims{
		JwtId: jwtId,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + config.Jwt.Time,
			Issuer:    config.Jwt.Issuer,
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtKey)
	if err != nil {
		return err.Error()
	} else {
		return token
	}
}

var ErrToken error = errors.New("token error")

func ParseToken(token string) (string, error) {
	temp, err := jwt.ParseWithClaims(token, &myClaims{}, func(token *jwt.Token) (interface{}, error) { return jwtKey, nil })
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return "", ErrToken //Token不正确
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				name := temp.Claims.(*myClaims).JwtId
				return name, ErrToken //Token已过期
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return "", ErrToken //Token无效
			} else {
				return "", ErrToken //这不是一个token
			}
		}

	}
	if temp != nil {
		if claims, ok := temp.Claims.(*myClaims); ok && temp.Valid {
			jwtId := claims.JwtId
			if jwtId == "" {
				return "", ErrToken
			}
			return jwtId, nil
		}
	}
	return "", ErrToken
}
