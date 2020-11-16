package model

type TeamAdd struct {
	Common
	Team TeamDefine `json:"team"`
}
