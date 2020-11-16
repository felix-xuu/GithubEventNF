package model

type PullRequestReviewDefine struct {
	User           Owner  `json:"user"`
	Body           string `json:"body"`
	State          string `json:"state"`
	HtmlUrl        string `json:"html_url"`
	PullRequestUrl string `json:"pull_request_url"`
}

type PullRequestReview struct {
	Common
	PullRequest PullRequestDefine       `json:"pull_request"`
	Review      PullRequestReviewDefine `json:"review"`
}
