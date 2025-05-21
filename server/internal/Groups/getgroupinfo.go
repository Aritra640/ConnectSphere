// TODO: make the validation, bind and uuid parser concurrent
package groups

import (
	"log"
	"net/http"

	"github.com/Aritra640/ConnectSphere/server/Database/db"
	"github.com/Aritra640/ConnectSphere/server/internal/utils"
	"github.com/labstack/echo/v4"
)

type GetGroupInfoParam struct {
	GroupID string `json:"group_id" validate:"required"`
}

// Get group details such as group name, group logo/image, restrictions, total number of members , won't get the list of the members in the group
func (g *GroupService) GetGroupInfoHandler(c echo.Context) error {
	var req GetGroupInfoParam
	err := c.Bind(&req)
	if err != nil {

		log.Println("Error: could not get proper request data in GetGroupInfoHandler: ", err)
		return c.JSON(404, "Invalid Request")
	}
	err = c.Validate(&req)
	if err != nil {

		log.Println("Error: could not get request data validated (validation error): ", err)
		return c.JSON(404, "Invalid Type")
	}

	guid, err := utils.ParseUUID(req.GroupID)
	if err != nil {
		log.Println("Error: Failed to convert groupid to uuid in GetGroupInfoHandler: ", err)
		return c.JSON(404, "Invalid Request")
	}

	grpCh := make(chan db.ChatGroup)
	defer close(grpCh)
	errCh := make(chan error)
	defer close(errCh)

	go func() {

		grpdata, err := g.Queries.GetGroupByID(c.Request().Context(), guid)
		if err != nil {
			errCh <- err
		}
		grpCh <- grpdata
	}()

	select {
	case err := <-errCh:
		log.Println(err)
		log.Println("Error: an error has occured in getting group information from database with the groupid: ", guid)
		return c.JSON(404, "Group Information not found!")

	case grp := <-grpCh:
		return c.JSON(http.StatusOK, map[string]interface{}{
			"id":                  grp.ID,
			"name":                grp.Name,
			"about":               grp.About,
			"ppic":                grp.Ppic.String,
			"required_permission": grp.RequiredPermission,
		})
	}
}
