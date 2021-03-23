package app

import (
	"github.com/IntroRestAPI/oauth-api/src/api/controllers/oauth"
	"github.com/IntroRestAPI/src/api/controllers/polo"
)

func mapUrls() {
	router.GET("/macro", polo.Polo)
	router.POST("/oauth/access_token", oauth.CreateAccessToken)
	router.GET("/oauth/access_token/:token_id", oauth.GetAccessToken)
}
