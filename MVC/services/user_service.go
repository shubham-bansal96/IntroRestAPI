package services

import (
	"github.com/IntroRestAPI/MVC/model"
	"github.com/IntroRestAPI/MVC/utils"
)

type usersService struct{

}

var(
	UsersService usersService
)

func(u usersService) GetUser(userid int)(*model.User, *utils.ApplicationError){
	return model.UserDao.GetUser(userid)
}