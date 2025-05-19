package users

import "github.com/Aritra640/ConnectSphere/server/Database/db"

type UserService struct {

	Queries *db.Queries
	Resentlink string
}


var UserSetup = &UserService{}
