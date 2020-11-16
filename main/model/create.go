package model

type Create struct {
	Common
	Ref          string `json:"ref"`
	RefType      string `json:"ref_type"`
	Description  string `json:"description"`
	MasterBranch string `json:"master_branch"`
}
