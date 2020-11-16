package model

type PageBuildDefine struct {
	Status string `json:"status"`
	Pusher Owner  `json:"pusher"`
	Error  struct {
		Message string `json:"message"`
	} `json:"error"`
}

type PageBuild struct {
	Common
	Build PageBuildDefine `json:"build"`
}
