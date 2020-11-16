package model

type OrgBlock struct {
	Common
	BlockedUser Owner `json:"blocked_user"`
}
