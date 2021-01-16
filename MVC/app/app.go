package app

import(
	"github.com/gin-gonic/gin"
	"github.com/IntroRestAPI/MVC/controllers"
)

var(
	router *gin.Engine
)

func init(){
	router = gin.Default()
}

func StartApp(){

	MapUrls()

	if err := router.Run(":4300"); err!=nil{
		panic(err)
	}
}

func MapUrls(){
	router.GET("/users/:user_id",controllers.GetUser)
}