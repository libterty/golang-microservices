package github_provider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"../../clients/restclient"
	"../../domains/github"
)

const (
	headerAuthorization = "Authorization"
	headerAuthorizationFormat = "token %s"

	urlCreateGithubRepo = "http://api.github.com/user/repos"
)

func getAuthorizationHeader(str string) string  {
	return fmt.Sprintf(headerAuthorizationFormat, str)
}

func CreateRepo(accessToken string, request github.CreateRepoRequest) (*github.CreateRepoResponse, *github.GithubErrorResponse)  {
	headers := http.Header{}
	headers.Set(headerAuthorization, getAuthorizationHeader(accessToken))
	response, err := restclient.Post(urlCreateGithubRepo, headers, request)
	if err != nil {
		log.Println(fmt.Sprintf("Error when trying to create github: %s", err.Error()))
		return nil, &github.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, &github.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message: "Invalid response body return",
		}
	}
	// after return of the following func the defer will be call and execute from defer stack
	defer response.Body.Close()

	if response.StatusCode > 299 {
		 var errResponse github.GithubErrorResponse
		 if err := json.Unmarshal(bytes, &errResponse); err != nil {
			 return nil, &github.GithubErrorResponse{
				 StatusCode: http.StatusInternalServerError,
				 Message: "Invalid json  body return",
			 }
		 }
		 errResponse.StatusCode = response.StatusCode
		 return nil, &errResponse
	}

	var result github.CreateRepoResponse
	if err := json.Unmarshal(bytes, &result); err != nil {
		log.Println(fmt.Sprintf("Error when unmarshal result: %s", err.Error()))
		return nil, &github.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message: "Invalid result response",
		}
	}

	return &result, nil
}