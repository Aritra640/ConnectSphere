package users

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/Aritra640/ConnectSphere/server/Database/db"
	"github.com/labstack/echo/v4"
)

type SetUserParam struct {
	PImage string `json:"pimage" validate:"required"`
	PBio   string `json:"pbio" validate:"required"`
}

// Add user info to the database with user id
func (u *UserService) AddUserInfoHandler(c echo.Context) error {

	userID := c.Request().Header.Get("userID")
	uid, err := strconv.Atoi(userID)
	if err != nil {
		log.Printf("Error: userid: %v not convertable or not found err: %v", userID, err)
		return c.JSON(500, "something went wrong")
	}

	var req SetUserParam
	err = c.Bind(&req)
	if err != nil {
		log.Println("Error: cannot get request in AddUserInfoHandler: ", err)
		return c.JSON(404, "Invalid Request")
	}
	err = c.Validate(&req)
	if err != nil {
		log.Println("Error: cannot validate request in AddUserInfoHandler: ", err)
		return c.JSON(404, "Invalid Type")
	}

	doneCh := make(chan bool)
	defer close(doneCh)

	go func() {
		_, err := u.Queries.AddUserInfo(c.Request().Context(), db.AddUserInfoParams{
			UserID: int32(uid),
			Pimage: sql.NullString{
				String: req.PImage,
				Valid:  true,
			},
			Pbio: sql.NullString{
				String: req.PBio,
				Valid:  true,
			},
		})

		if err != nil {
			log.Println("Error: could not add to users_info table in the database with the folowing error: ", err)
			doneCh <- false
		}
		doneCh <- true
	}()

	done := <-doneCh
	if done {
		return c.JSON(http.StatusOK, "Userinfo added")
	}
	log.Println("Used Info cannot be added with the following user id: ", uid)
	return c.JSON(500, "Something went wrong")
}
