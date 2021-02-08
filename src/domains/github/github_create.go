package github

/**
* @ struct
{
   "name": "golang-microservices",
   "description": "golang-microservices-practice",
   "homepage": "http://github.com",
   "private": false,
   "has_issues": true,
   "has_projects": true,
   "has_wiki": true
}
*/

type CreateRepoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
	Private     bool   `json:"private"`
	HasIssues   bool   `json:"has_issues"`
	HasProjects bool   `json:"has_projects"`
	HasWiki     bool   `json:"has_wiki"`
}
