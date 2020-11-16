package model

type GollumDefine struct {
	PageName string `json:"page_name"`
	Title    string `json:"title"`
	Summary  string `json:"summary"`
	Action   string `json:"action"`
	HtmlUrl  string `json:"html_url"`
}

type Gollum struct {
	Common
	Pages []GollumDefine `json:"pages"`
}
