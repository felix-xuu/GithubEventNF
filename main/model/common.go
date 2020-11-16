package model

type Common struct {
	Action       string             `json:"action"`
	Sender       Owner              `json:"sender"`
	Installation InstallationDefine `json:"installation"`
	Organization OrganizationDefine `json:"organization"`
	Repository   RepositoryDefine   `json:"repository"`
}
