package model

type DeploymentDefine struct {
	Ref                 string            `json:"ref"`
	Task                string            `json:"task"`
	Payload             map[string]string `json:"payload"`
	OriginalEnvironment string            `json:"original_environment"`
	Environment         string            `json:"environment"`
	Description         string            `json:"description"`
	Creator             Owner             `json:"creator"`
}

type Deployment struct {
	Common
	Deployment DeploymentDefine `json:"deployment"`
}
