package model

type CommitCommentDefine struct {
	HtmlUrl string `json:"html_url"`
	User    Owner  `json:"user"`
	Body    string `json:"body"`
}

type CommitComment struct {
	Common
	Comment CommitCommentDefine `json:"comment"`
}
