package model

type ReleaseDefine struct {
	HtmlUrl string `json:"html_url"`
	TagName string `json:"tag_name"`
	Name    string `json:"name"`
	Author  Owner  `json:"author"`
	Body    string `json:"body"`
}

type Release struct {
	Common
	Release ReleaseDefine `json:"release"`
}
