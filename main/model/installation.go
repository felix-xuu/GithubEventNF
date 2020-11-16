package model

type InstallationDefine struct {
	Account             Owner  `json:"account"`
	RepositorySelection string `json:"repository_selection"`
	HtmlUrl             string `json:"html_url"`
	SingleFileName      string `json:"single_file_name"`
}

type Installation struct {
	Common
	Installation InstallationDefine `json:"installation"`
	Repositories []RepositoryDefine `json:"repositories"`
}
