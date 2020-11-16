package model

type IssueDefine struct {
	HtmlUrl   string          `json:"html_url"`
	Title     string          `json:"title"`
	State     string          `json:"state"`
	Locked    bool            `json:"locked"`
	Comments  int32           `json:"comments"`
	Body      string          `json:"body"`
	User      Owner           `json:"user"`
	Labels    []LabelDefine   `json:"labels"`
	Milestone MilestoneDefine `json:"milestone"`
	Assignee  Owner           `json:"assignee"`
	Assignees []Owner         `json:"assignees"`
}

type Issues struct {
	Common
	Issue IssueDefine `json:"issue"`
}
