package auth

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc { 
  return func(c echo.Context) error {

    jwtToken := c.Request().Header.Get("JWT")
    if jwtToken == "" {
      log.Println("Error: could not get auth header")
      return c.JSON(http.StatusBadRequest , map[string]string {
        "error": "could not get jwt header!",
      })
    }

    log.Println("JWT: " , jwtToken)

    user_id,err := VerifyToken(jwtToken)
    if err != nil {
      log.Println("Error in verifying token")
      return c.JSON(http.StatusBadRequest , map[string]string{
        "error" : "could not verify token",
      })
    }

    c.Request().Header.Add("UserID" , strconv.Itoa(user_id))
    return next(c)
  }
}
