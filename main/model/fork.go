package model

type Fork struct {
	Common
	Forkee RepositoryDefine `json:"forkee"`
}
