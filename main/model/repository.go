package model

type RepositoryDefine struct {
	Name            string `json:"name"`
	FullName        string `json:"full_name"`
	Owner           Owner  `json:"owner"`
	HtmlUrl         string `json:"html_url"`
	Description     string `json:"description"`
	StargazersCount int32  `json:"stargazers_count"`
	WatchersCount   int32  `json:"watchers_count"`
	ForksCount      int32  `json:"forks_count"`
	OpenIssuesCount int32  `json:"open_issues_count"`
	Forks           int32  `json:"forks"`
	OpenIssues      int32  `json:"open_issues"`
	Watchers        int32  `json:"watchers"`
	DefaultBranch   string `json:"default_branch"`
	Language        string `json:"language"`
}

type Repository struct {
	Common
}
