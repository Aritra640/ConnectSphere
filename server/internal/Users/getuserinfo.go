package users

import (
	"log"
	"strconv"

	"github.com/Aritra640/ConnectSphere/server/Database/db"
	"github.com/labstack/echo/v4"
)

type UserInfoReturn struct {
	PImage string `json:"pimage"`
	PBio   string `json:"pbio"`
}

// Get user details (profile image and profile bio) from the database through user id
func (u *UserService) GetUserInfo(c echo.Context) error {

	userID := c.Request().Header.Get("userID")
	log.Println("Request received to get user info with user id: ", userID)

	uid, err := strconv.Atoi(userID)
	if err != nil {
		log.Println("Userid could not be converted to integer in GetUserInfo: ", err)
		return c.JSON(404, "Someting went wrong")
	}

	uifCh := make(chan db.UsersInfo)
	defer close(uifCh)
	errChan := make(chan error)

	go func() {
		userdata, err := u.Queries.GetUserInfo(c.Request().Context(), int32(uid))
		if err != nil {
			log.Println("Error: Could not get userdata from database: ", err)
			errChan <- err
		}
		uifCh <- userdata
	}()

	select {
	case <-errChan:
		return c.JSON(500, "Somehting Went wrong")

	case user := <-uifCh:
		ret := UserInfoReturn{
			PImage: user.Pimage.String,
			PBio:   user.Pbio.String,
		}

		return c.JSON(200, ret)
	}
}
