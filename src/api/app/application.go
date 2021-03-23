package app

import (
	"github.com/IntroRestAPI/src/api/controllers/polo"
	"github.com/IntroRestAPI/src/api/controllers/repositories"
	"github.com/IntroRestAPI/src/api/log"
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

func StartApp() {
	log.Info("about to map URLs", "step:1", "status:pending")
	MapUrls()
	log.Info("URLs successfully Mapped", "step:2", "status:success")

	if err := router.Run(":4300"); err != nil {
		panic(err)
	}
}

func MapUrls() {
	router.GET("/macro", polo.Polo)
	router.POST("/repositories", repositories.CreateRepo)
}
