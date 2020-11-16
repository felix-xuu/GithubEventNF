package model

type DeployKeyDefine struct {
	Key      string `json:"key"`
	Title    string `json:"title"`
	Verified string `json:"verified"`
	ReadOnly string `json:"read_only"`
}

type DeployKey struct {
	Common
	Key DeployKeyDefine `json:"key"`
}
