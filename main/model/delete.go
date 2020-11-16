package model

type Delete struct {
	Common
	Ref     string `json:"ref"`
	RefType string `json:"ref_type"`
}
