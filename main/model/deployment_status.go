package model

type DeploymentStatusDefine struct {
	State       string `json:"state"`
	Creator     Owner  `json:"creator"`
	Description string `json:"description"`
	Environment string `json:"environment"`
	TargetUrl   string `json:"target_url"`
}

type DeploymentStatus struct {
	Common
	DeploymentStatus DeploymentStatusDefine `json:"deployment_status"`
	Deployment       DeploymentDefine       `json:"deployment"`
}
