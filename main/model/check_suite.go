package model

type CheckSuiteDefine struct {
	HeadBranch string `json:"head_branch"`
	Status     string `json:"status"`
	Conclusion string `json:"conclusion"`
}

type CheckSuite struct {
	Common
	CheckSuite CheckSuiteDefine `json:"check_suite"`
}
