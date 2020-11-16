package model

type Member struct {
	Common
	Member Owner `json:"member"`
}
