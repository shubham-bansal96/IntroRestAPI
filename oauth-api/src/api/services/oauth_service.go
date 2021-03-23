package services

import (
	"time"

	"github.com/IntroRestAPI/oauth-api/src/api/domain/oauth"
	"github.com/IntroRestAPI/src/api/utils/errors"
)

type oauthService struct{}

type oauthServiceInterface interface {
	CreateAccessToken(request oauth.AccessTokenRequest) (*oauth.AccessToken, errors.ApiError)
	GetAccessToken(accessToken string) (*oauth.AccessToken, errors.ApiError)
}

var (
	OauthService oauthServiceInterface
)

func init() {
	OauthService = &oauthService{}
}

func (s *oauthService) CreateAccessToken(request oauth.AccessTokenRequest) (*oauth.AccessToken, errors.ApiError) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	user, err := oauth.GetUserByUsernameAndPassword(request.Username, request.Password)
	if err != nil {
		return nil, err
	}

	token := oauth.AccessToken{
		UserId:  user.Id,
		Expires: int(time.Now().UTC().Add(24 * time.Hour).Unix()),
	}

	if err := token.Save(); err != nil {
		return nil, err
	}

	return &token, nil
}

func (s *oauthService) GetAccessToken(accessToken string) (*oauth.AccessToken, errors.ApiError) {
	token, err := oauth.GetAccessTokenByToken(accessToken)

	if err != nil {
		return nil, err
	}

	if token.IsExpired() {
		return nil, errors.NewNotFoundError("no access token found with given parameters")
	}
	return token, err
}
