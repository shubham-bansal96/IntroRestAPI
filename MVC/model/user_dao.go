package model

import(
	"fmt"
	"github.com/IntroRestAPI/MVC/utils"
	"net/http"
)

var (
	users = map[int]*User{
		123: {Id:1, FirstName:"shubham", LastName:"bansal",Email:"shubham123@gmail.com"},
		321: {Id:2, FirstName:"ashish", LastName:"bhatt",Email:"ashish321@gmail.com"},
	}

	UserDao userDao
)

type userDao struct{

}

func(u userDao) GetUser(userid int)(*User, *utils.ApplicationError){
	if user := users[userid]; user != nil{
		return user,nil		
	}
	return nil, &utils.ApplicationError{
		Message: fmt.Sprintf("user %v was not found", userid),
		StatusCode: http.StatusNotFound,
		Code: "not found",
	}
}