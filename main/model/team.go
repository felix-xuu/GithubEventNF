package model

type TeamDefine struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Privacy     string `json:"privacy"`
	HtmlUrl     string `json:"html_url"`
	Permission  string `json:"permission"`
}

type Team struct {
	Common
	Team TeamDefine `json:"team"`
}
