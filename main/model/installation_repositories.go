package model

type InstallationRepositories struct {
	Common
	RepositorySelection string             `json:"repository_selection"`
	RepositoriesAdded   []RepositoryDefine `json:"repositories_added"`
	RepositoriesRemoved []RepositoryDefine `json:"repositories_removed"`
}
