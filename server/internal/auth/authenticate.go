package auth

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/Aritra640/ConnectSphere/server/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

type MyCustomClaims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

func CreateAuthToken(UserID int) (string, error) {

	claims := MyCustomClaims{
		UserID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	ss, err := token.SignedString(config.App.JWT)
	if err != nil {
		log.Println("could not create jwt token")
		return "", err
	}

	return ss, nil
}

func VerifyToken(tokenString string) (int, error) {

	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		//Validate the signing method (HMAC)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method : %v", token.Header["alg"])
		}
		return config.App.JWT, nil
	})

	if err != nil {
		log.Println("Error in parsing token: ", err.Error())
		return -1, err
	}

	claims, ok := token.Claims.(*MyCustomClaims)
	if !ok || !token.Valid {
		log.Println("invalid token or claims")
		return -1, errors.New("error cannot verify token")
	}

	log.Println("Token verified successfully for user-id: ", claims.UserID)
	return claims.UserID, nil
}

// CheckExpiry returns false if the token is valid but expired , returns true otherwise
func CheckExpiry(tokenString string) (int, bool) {

	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		//Validate the signing method (HMAC)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method : %v", token.Header["alg"])
		}
		return config.App.JWT, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			log.Println("Token is expired")
			claims,ok := token.Claims.(*MyCustomClaims)
      if !ok {
        log.Println("claims invalid101")
        return -1 , true
      }

      return claims.UserID , false
		}
	}

	log.Println("Token parsing failed for another reason: ", err)
	return -1,true
}
