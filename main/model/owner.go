package model

type Owner struct {
	Login                    string `json:"login"`
	AvatarUrl                string `json:"avatar_url"`
	HtmlUrl                  string `json:"html_url"`
	OrganizationBillingEmail string `json:"organization_billing_email"`
}
