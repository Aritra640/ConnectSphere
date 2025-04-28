package ws

import (
	"log"
	"net/http"

	"github.com/Aritra640/ConnectSphere/server/Database/db"
	"github.com/labstack/echo/v4"
)

type GetFriends struct {
	UserID int `json:"user_id"`
}

func (pcs *PersonalChatService) GetAllFriendsHandler(c echo.Context) error {

	var req GetFriends
	if err := c.Bind(&req); err != nil {
		log.Println("Error: could not get user_id in GetAllFriendsHandler : ", err)
		return c.JSON(http.StatusBadRequest, "Invalid Request")
	}

	friendsCh := make(chan []db.User)
	errCh := make(chan error)

	go func() {

		friends, err := pcs.Queries.GetUserFriends(c.Request().Context(), int32(req.UserID))
		if err != nil {
			errCh <- err
		}
		friendsCh <- friends
	}()
	var err error
	var friends []db.User

	select {
	case err = <-errCh:
		log.Println("Error: cannot find friends with user id: ", req.UserID)
		log.Println("Error: cannot get friends in user database: ", err)
		return c.JSON(http.StatusConflict, "Something went wrong")

	case friends = <-friendsCh:
		res := JSONfriendsConverter(friends)
		log.Println("Successfully found users friends with id: ", req.UserID)
		return c.JSON(http.StatusOK, res)
	}
}

type JSONfriends struct {
	UserID    int    `json:"user_id"`
	UserName  string `json:"user_name"`
	UserEmail string `json:"user_email"`
}

func JSONfriendsConverter(friends []db.User) []JSONfriends {

	resCh := make(chan []JSONfriends)

	go func() {

		res := make([]JSONfriends, 0)
		for _, friend := range friends {

			newFriend := JSONfriends{
				UserID:    int(friend.ID),
				UserName:  friend.Username,
				UserEmail: friend.Email,
			}

			res = append(res, newFriend)
		}

		resCh <- res
	}()

	res := <-resCh
	return res
}

func SampleGetAllFriendsHandler(c echo.Context) error {

  temp := 'a'
	var res []JSONfriends

	for i := 0; i < 3; i++ {
		iter := JSONfriends{
			UserName:  string(temp),
			UserID:    i,
			UserEmail: string(temp),
		}

		res = append(res, iter)

		temp++
	}

	return c.JSON(http.StatusOK, res)
}
