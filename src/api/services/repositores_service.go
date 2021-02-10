package services

import (
	"strings"

	"../config"
	"../domains/github"
	"../domains/repositories"
	"../providers/github_provider"
	"../utils/errors"
)

type reposService struct{}

var (
	RepositoryService reposService
)

func (s *reposService) CreateRepo(req repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiErrorInterface) {
	req.Name = strings.TrimSpace(req.Name)
	if req.Name == "" {
		return nil, errors.NewBadRequestApiError("Empty name input")
	}

	request := github.CreateRepoRequest{
		Name:        req.Name,
		Description: req.Description,
		Private:     false,
		HasWiki:     true,
		HasProjects: true,
		HasIssues:   true,
	}

	response, err := github_provider.CreateRepo(config.GetGithubAccessToken(), request)
	if err != nil {
		return nil, errors.NewApiError(err.StatusCode, err.Message)
	}

	result := repositories.CreateRepoResponse{
		Id:    response.Id,
		Name:  response.Name,
		Owner: response.Owner.Login,
	}

	return &result, nil
}
