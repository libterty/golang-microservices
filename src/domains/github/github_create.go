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

type CreateRepoResponse struct {
	Id          int64        `json:"id"`
	Name        string       `json:"name"`
	FullName    string       `json:"full_name"`
	Owner       RepoOwner    `json:"owner"`
	Permissions RepoMissions `json:"permissions"`
}

type RepoOwner struct {
	Id      int64  `json:"id"`
	Login   string `json:"login"`
	Url     string `json:"url"`
	HtmlUrl string `json:"html_url"`
}

type RepoMissions struct {
	IsAdmin  bool `json:"is_admin"`
	HasPull  bool `json:"has_pull"`
	HashPush bool `json:"hash_push"`
}
