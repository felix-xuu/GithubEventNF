package model

type CheckRunDefine struct {
	HtmlUrl    string `json:"html_url"`
	DetailsUrl string `json:"details_url"`
	Status     string `json:"status"`
	Name       string `json:"name"`
	Conclusion string `json:"conclusion"`
}

type CheckRun struct {
	Common
	CheckRun CheckRunDefine `json:"check_run"`
}
