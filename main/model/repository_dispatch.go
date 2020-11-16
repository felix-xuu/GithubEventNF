package model

type RepositoryDispatch struct {
	Common
	Action string `json:"action"`
	Branch string `json:"branch"`
}
