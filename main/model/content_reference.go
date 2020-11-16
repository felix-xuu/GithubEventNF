package model

type ContentReferenceDefine struct {
	Reference string `json:"reference"`
}

type ContentReference struct {
	Common
	ContentReference ContentReferenceDefine `json:"content_reference"`
}
