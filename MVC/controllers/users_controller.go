package controllers

import (
	"net/http"
	"strconv"
	"github.com/IntroRestAPI/MVC/services"
	"github.com/IntroRestAPI/MVC/utils"
	"github.com/gin-gonic/gin"
)


func GetUser(c *gin.Context){
	userId,err:= strconv.ParseInt(c.Param("user_id"),10,64)

	if err != nil{
		apiErr:= utils.ApplicationError{
			Message:"user_id must be a number",
			StatusCode:http.StatusBadRequest,
			Code:"bad_request",
		}
		utils.Respond(c, apiErr.StatusCode, apiErr)
		return
	}

	
	user, apiErr := services.UsersService.GetUser(int(userId))
	if apiErr != nil{
		utils.Respond(c, apiErr.StatusCode, apiErr)
		return
	}

	//return user to client
	utils.Respond(c, http.StatusOK, user)
}