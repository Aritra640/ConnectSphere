package users

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/Aritra640/ConnectSphere/server/Database/db"
	"github.com/labstack/echo/v4"
)

type UpdateUserInfoImageParam struct {
	PImage string `json:"pimage"`
}

// Update user profile image
func (u *UserService) UpdateUserInfoImageHandler(c echo.Context) error {

	userID := c.Request().Header.Get("userID")
	if userID == "" {
		log.Println("Error: Userid could not be received in UpdateUserInfoImageHandler")
		return c.JSON(500, "Internal error")
	}
	uid, err := strconv.Atoi(userID)
	if err != nil {
		log.Println("Error: could not get user id in proper format in UpdateUserInfoImageHandler , userID: ", userID)
		return c.JSON(500, "Something Went Wrong")
	}

	var req UpdateUserInfoImageParam
	err = c.Bind(&req)
	if err != nil {
		log.Println("Error: Couldnot get request data: ", err)
		return c.JSON(404, "Imvalid Request")
	}

	errCh := make(chan error)
	defer close(errCh)
	go func() {

		err := u.Queries.UpdateUserInfoImage(c.Request().Context(), db.UpdateUserInfoImageParams{
			UserID: int32(uid),
			Pimage: sql.NullString{
				String: req.PImage,
				Valid:  true,
			},
		})

		errCh <- err
	}()

	err = <-errCh
	if err != nil {
		log.Println("Error: Couldnot update user info data(image) in the database: ", err)
		return c.JSON(http.StatusInternalServerError, "Something went wrong")
	}

	return c.JSON(http.StatusOK, "User profile image successfully updated")
}

type UpdateUserInfoBioParam struct {
	PBio string `json:"pbio"`
}

// Update user profile bio
func (u *UserService) UpdateUserInfoBioHandler(c echo.Context) error {

	userID := c.Request().Header.Get("userID")
	if userID == "" {
		log.Println("Error: Userid could not be received in UpdateUserInfoImageHandler")
		return c.JSON(500, "Internal error")
	}
	uid, err := strconv.Atoi(userID)
	if err != nil {
		log.Println("Error: could not get user id in proper format in UpdateUserInfoImageHandler , userID: ", userID)
		return c.JSON(500, "Something Went Wrong")
	}

	var req UpdateUserInfoBioParam
	err = c.Bind(&req)
	if err != nil {
		log.Println("Error: Invalid request in UpdateUserInfoBioHandler: " , err)
		return c.JSON(404 , "Invalid Request")
	}
	
	errChan := make(chan error)
	defer close(errChan)

	go func() {
		
		err := u.Queries.UpdateUserInfoBio(c.Request().Context() , db.UpdateUserInfoBioParams{
			UserID: int32(uid),
			Pbio: sql.NullString{
				String: req.PBio,
				Valid: true,
			},
		})
		errChan <- err
	}()

	err = <-errChan 
	if err != nil {
		log.Println("Error: Could not update profile bio in database in UpdateUserInfoBioHandler: " , err)
		return c.JSON(500 , "Something went wrong")
	}

	return c.JSON(200 , "User profile bio updated!")
}
