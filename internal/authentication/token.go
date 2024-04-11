package authentication

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/rcarvalho-pb/todo-app-go/internal/infra/configs"
)

func CreateToken(userId uint32) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"authorized": true,
		"exp":        time.Now().Add(time.Hour * 1).Unix(),
		"userId":     userId,
	})

	return token.SignedString([]byte(configs.EnvConfigs.SecretKey))
}

func ValidateToken(r *http.Request) error {
	token, err := extractToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.MapClaims); !ok {
		return errors.New("invalid token")
	}

	return nil
}

func getValidationKey(token *jwt.Token) (any, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}

	return []byte(configs.EnvConfigs.SecretKey), nil
}

func extractToken(r *http.Request) (*jwt.Token, error) {
	token := r.Header.Get("Authorization")
	if len(strings.Split(token, " ")) == 2 {
		token, err := jwt.Parse(strings.Split(token, " ")[1], getValidationKey)
		if err != nil {
			return nil, err
		}

		return token, nil
	}

	return nil, errors.New("invalid token")
}

func ExtractUserId(r *http.Request) (uint32, error) {
	token, err := extractToken(r)
	if err != nil {
		return 0, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId, err := strconv.ParseUint(fmt.Sprintf("%v", claims["userId"]), 10, 32)
		if err != nil {
			return 0, err
		}
		return uint32(userId), nil
	}
	return 0, errors.New("invalid token")
}
