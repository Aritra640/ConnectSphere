package users

import "github.com/labstack/echo/v4"


func (u *UserService) FriendInfoHandler(c echo.Context) error {

	return c.JSON(200 , "friends info")
}
