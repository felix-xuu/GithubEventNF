package model

type LabelDefine struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

type Label struct {
	Common
	Label LabelDefine `json:"label"`
}
