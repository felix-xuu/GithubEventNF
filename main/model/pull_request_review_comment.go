package model

type PullRequestReviewCommentDefine struct {
	User    Owner  `json:"user"`
	Path    string `json:"path"`
	Body    string `json:"body"`
	HtmlUrl string `json:"html_url"`
}

type PullRequestReviewComment struct {
	Common
	PullRequest PullRequestDefine              `json:"pull_request"`
	Comment     PullRequestReviewCommentDefine `json:"comment"`
}
