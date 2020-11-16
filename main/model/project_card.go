package model

type ProjectCardDefine struct {
	Note     string `json:"note"`
	Archived bool   `json:"archived"`
	Creator  Owner  `json:"creator"`
}

type ProjectCard struct {
	Common
	ProjectCard ProjectCardDefine `json:"project_card"`
}
