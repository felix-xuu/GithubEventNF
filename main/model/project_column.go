package model

type ProjectColumnDefine struct {
	Name string `json:"name"`
}

type ProjectColumn struct {
	Common
	ProjectColumn ProjectColumnDefine `json:"project_column"`
}
