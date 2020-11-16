package model

type CodeScanningAlertDefine struct {
	HtmlUrl string `json:"html_url"`
	State   string `json:"state"`
}

type CodeScanningAlert struct {
	Common
	Alert CodeScanningAlertDefine `json:"alert"`
}
