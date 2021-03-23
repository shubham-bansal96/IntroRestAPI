package services

import (
	"strings"

	"github.com/IntroRestAPI/src/api/config"
	"github.com/IntroRestAPI/src/api/domain/github"
	"github.com/IntroRestAPI/src/api/providers/github_provider"

	"github.com/IntroRestAPI/src/api/domain/repositories"
	"github.com/IntroRestAPI/src/api/utils/errors"
)

type repoService struct{}

type reposServiceInterface interface {
	CreateRepo(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
}

var (
	RepositoryService reposServiceInterface
)

func init() {
	RepositoryService = &repoService{}
}

func (s *repoService) CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
	if input.Name = strings.TrimSpace(input.Name); input.Name == "" {
		return nil, errors.NewBadRequestError("invalid repository name")
	}

	request := github.CreateRepoRequest{
		Name:        input.Name,
		Description: input.Description,
		Private:     false,
	}

	response, err := github_provider.CreateRepo(config.GetGithubAccessToken(), request)
	if err != nil {
		return nil, errors.NewApiError(err.StatusCode, err.Message)
	}

	result := &repositories.CreateRepoResponse{
		Id:    response.Id,
		Name:  response.Name,
		Owner: response.Owner.Login,
	}
	return result, nil
}
