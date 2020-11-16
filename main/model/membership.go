package model

type Membership struct {
	Common
	Scope           string     `json:"scope"`
	Member          Owner      `json:"member"`
	Team            TeamDefine `json:"team"`
	OrganizationUrl string     `json:"organization_url"`
	State           string     `json:"state"`
	Role            string     `json:"role"`
	User            Owner      `json:"user"`
}
