package github

type CreateRepoRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Homepage    string `json:"homepage,omitempty"`
	Private     bool   `json:"private,omitempty"`
	Hasissues   bool   `json:"has_issues,omitempty"`
	Hasprojects bool   `json:"has_projects,omitempty"`
	Haswiki     bool   `json:"has_wiki,omitempty"`
}

type CreateRepoResponse struct {
	Id          int             `json: "id, omitempty"`
	Name        string          `json: "id, omitempty"`
	FullName    string          `json: "id, omitempty"`
	Owner       RepoOwner       `json: "owner, omitempty"`
	Permissions RepoPermissions `json: "permissions, omitempty"`
}

type RepoOwner struct {
	Id      int    `json: "id, omitempty"`
	Login   string `json: "login, omitempty"`
	Url     string `json: "url, omitempty"`
	HtmlUrl string `json: "html_url, omitempty"`
}

type RepoPermissions struct {
	IsAdmin bool `json: "admin, omitempty"`
	HasPull bool `json: "pull, omitempty"`
	HasPush bool `json: "push, omitempty"`
}
