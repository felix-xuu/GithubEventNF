package model

type MilestoneDefine struct {
	HtmlUrl      string `json:"html_url"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Creator      Owner  `json:"creator"`
	OpenIssues   int32  `json:"open_issues"`
	ClosedIssues int32  `json:"closed_issues"`
	State        string `json:"state"`
}

type Milestone struct {
	Common
	Milestone MilestoneDefine `json:"milestone"`
}
