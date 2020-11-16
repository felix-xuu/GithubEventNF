package model

type ProjectDefine struct {
	HtmlUrl string `json:"html_url"`
	Name    string `json:"name"`
	Body    string `json:"body"`
	State   string `json:"state"`
	Creator Owner  `json:"creator"`
}

type Project struct {
	Common
	Project ProjectDefine `json:"project"`
}
