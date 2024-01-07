package app

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/mhmdiamd/go-restapi-future-store/exception"
	"github.com/mhmdiamd/go-restapi-future-store/helper"
)

type JWTPayload struct {
  ID uuid.UUID `json:"id"`
  Email string `json:"email"`
  Role string `json:"role"`
}

func GenerateJWT (payload JWTPayload) (string, error) {
  token := jwt.New(jwt.SigningMethodHS256)
  claims := token.Claims.(jwt.MapClaims)

  claims["authorized"] = true
  claims["ID"] = payload.ID
  claims["Role"] = payload.Role
  claims["Email"] = payload.Email
  claims["exp"] = time.Now().Add(time.Minute * 60).Unix()

  tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

  helper.PanicIfError(err)

  return "Bearer " + tokenString, nil

}

func ValidateToken(token string) (JWTPayload, error) {

  // Split the value and only take the token
  newToken := strings.Split(token, " ")[1]

  payload := JWTPayload{}

	parsedToken, _ := jwt.Parse(newToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("There was an error in parsing")
		}
    return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if parsedToken == nil {
    return payload, errors.New("Invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return payload, errors.New("Couldn't parse claims")
	}

	exp := claims["exp"].(float64)
	if int64(exp) < time.Now().Local().Unix() {
	  panic(exception.NewForbiddenError("Token Expired"))
  }

  payload = JWTPayload{
    ID : uuid.MustParse(claims["ID"].(string)),
    Email : claims["Email"].(string),
    Role : claims["Role"].(string),
  }

	return payload, nil
}

func TokenParsedJWT(token string) (JWTPayload, error) {

  fmt.Println(os.Getenv("JWT_SECRET"))
  parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error ) {
    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
      return nil, errors.New("Error when parse token")
    }

    return []byte(os.Getenv("JWT_SECRET")), nil
  })

  var payload = JWTPayload{}

  if err != nil {
    return payload, errors.New("Token is expired")
  }

  tokenClaims, ok := parsedToken.Claims.(jwt.MapClaims)

  if !ok {
    return payload, errors.New("Couldn't parse token")
  }

  payload = JWTPayload{
    ID : uuid.MustParse(tokenClaims["ID"].(string)),
    Email : tokenClaims["Email"].(string) ,
    Role : tokenClaims["Role"].(string),
  }

  return payload, nil
}







