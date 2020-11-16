package model

type PullRequestDefine struct {
	HtmlUrl  string `json:"html_url"`
	DiffUrl  string `json:"diff_url"`
	PatchUrl string `json:"patch_url"`
	State    string `json:"state"`
	Title    string `json:"title"`
	Locked   bool   `json:"locked"`
	User     Owner  `json:"user"`
	Body     string `json:"body"`
}

type PullRequest struct {
	Common
	PullRequest PullRequestDefine `json:"pull_request"`
}
