package lib

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/hanherb/mini-wallet/src/config"
)

func CreateAccessToken(id uuid.UUID) (accessToken string, err error) {
	atClaims := jwt.MapClaims{}
	atClaims["id"] = id
	atClaims["iat"] = time.Now().Unix()
	atClaims["exp"] = time.Now().Add(time.Hour * 24 * 30).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	accessToken, err = at.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return
	}

	return
}

func ExtractToken(r *http.Request) (*jwtPayload, error) {
	token := r.Header.Get("Authorization")

	strArr := strings.Split(token, " ")
	if len(strArr) != 2 {
		return nil, errors.New("invalid/empty token")
	}

	payload, err := VerifyAccessToken(&strArr[1])

	if err != nil {
		return nil, err
	}

	return payload, nil
}

type jwtPayload struct {
	ID string `json:"id"`
}

func VerifyAccessToken(tokenString *string) (payload *jwtPayload, err error) {
	token, err := jwt.Parse(*tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.JWTSecretKey), nil
	})

	payload = new(jwtPayload)
	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		payload.ID = claims["id"].(string)
	} else {
		return payload, err
	}

	if err != nil {
		return payload, err
	}

	return payload, nil
}

func TokenCustomerId(c *gin.Context) (customerId uuid.UUID, err error) {
	jwtId := c.MustGet("jwtId").(string)
	customerId, err = uuid.FromString(jwtId)
	if err != nil {
		return
	}

	return
}
