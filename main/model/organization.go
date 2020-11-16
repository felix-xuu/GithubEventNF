package model

type OrganizationDefine struct {
	Login       string `json:"login"`
	AvatarUrl   string `json:"avatar_url"`
	Description string `json:"description"`
}

type Organization struct {
	Common
	Membership Membership `json:"membership"`
}
