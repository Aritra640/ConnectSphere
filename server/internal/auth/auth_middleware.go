//TODO: add refresh token redirect 
package auth

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (as *AuthService) AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc { 
  return func(c echo.Context) error {

    jwtToken := c.Request().Header.Get("access_token")
    if jwtToken == "" {
      log.Println("Error: could not get auth header")
      return c.JSON(http.StatusBadRequest , map[string]string {
        "error": "could not get jwt header!",
      })
    }

    log.Println("JWT: " , jwtToken)

    user_id,err := VerifyToken(jwtToken)
    if err != nil {

      returnCh := make(chan checkExpiryReturn)

      go func() {
        user_id,check := CheckExpiry(jwtToken)
        returnCh <- checkExpiryReturn{
          userID: user_id,
          flag: check,
        }
      }()
      ret := <-returnCh

      if !ret.flag  {
        log.Println("Token has expired ,verifying refresh token")
        
        refresh_token := c.Request().Header.Get("refresh_token")
        if refresh_token == "" {
          return c.JSON(http.StatusUnauthorized , map[string]string {
            "error": "missing refresh token",
          })
        }

        //Verify refresh token 
        _,err := as.Rts.VerifyRefreshToken(c.Request().Context() , ret.userID , refresh_token)
        if err != nil {
          log.Println("Error: Invalid refresh token: " , err)
          return c.JSON(http.StatusUnauthorized , "Invalid Refresh Token, please signin again")
        }

      }
      
      log.Println("Error in verifying token")
      return c.JSON(http.StatusBadRequest , map[string]string{
        "error" : "could not verify token",
      })
    }

    c.Request().Header.Add("UserID" , strconv.Itoa(user_id))
    return next(c)
  }
}


type checkExpiryReturn struct {

  userID int 
  flag   bool
}
