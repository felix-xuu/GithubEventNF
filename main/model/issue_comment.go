package model

type IssueCommentDefine struct {
	HtmlUrl string `json:"html_url"`
	User    Owner  `json:"user"`
	Body    string `json:"body"`
}

type IssueComment struct {
	Common
	Issue   IssueDefine        `json:"issue"`
	Comment IssueCommentDefine `json:"comment"`
}
