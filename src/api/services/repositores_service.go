package services

import (
	"../config"
	"../domains/github"
	"../domains/repositories"
	"../providers/github_provider"
	"../utils/errors"
	"net/http"
	"sync"
)

type reposService struct{}

type reposServiceInterface interface {
	CreateRepo(req repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiErrorInterface)
	CreateRepos(req []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.ApiErrorInterface)
}

var (
	RepositoryService reposServiceInterface
)

func init() {
	RepositoryService = &reposService{}
}

func (s *reposService) CreateRepo(req repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiErrorInterface) {
	if err := req.Validate(); err != nil {
		return nil, err
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

func (s *reposService) CreateRepos(req []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.ApiErrorInterface) {
	input := make(chan repositories.CreateRepositoriesResult)
	output := make(chan repositories.CreateReposResponse)
	defer close(output)

	var wq sync.WaitGroup
	go s.HandleRepoResults(&wq, input, output)

	for _, cur := range req {
		wq.Add(1)
		go s.CreateRepoConCurrent(cur, input)
	}

	wq.Wait()
	close(input)

	result := <-output

	successCreations := 0
	for _, current := range result.Results {
		if current.Response != nil {
			successCreations++
		}
	}
	if successCreations == 0 {
		result.StatusCode = result.Results[0].Error.Status()
	} else if successCreations == len(req) {
		result.StatusCode = http.StatusCreated
	} else {
		result.StatusCode = http.StatusPartialContent
	}
	return result, nil
}

func (s *reposService) HandleRepoResults(wq *sync.WaitGroup, input chan repositories.CreateRepositoriesResult, output chan repositories.CreateReposResponse) {
	var results repositories.CreateReposResponse

	for incomeEvent := range input {
		repoResult := repositories.CreateRepositoriesResult{
			Response: incomeEvent.Response,
			Error:    incomeEvent.Error,
		}
		results.Results = append(results.Results, repoResult)

		wq.Done()
	}
	output <- results
}

func (s *reposService) CreateRepoConCurrent(input repositories.CreateRepoRequest, out chan repositories.CreateRepositoriesResult) {
	if err := input.Validate(); err != nil {
		out <- repositories.CreateRepositoriesResult{Error: err}
		return
	}

	result, err := s.CreateRepo(input)
	if err != nil {
		out <- repositories.CreateRepositoriesResult{
			Error: err,
		}
		return
	}

	out <- repositories.CreateRepositoriesResult{
		Response: result,
	}
}
