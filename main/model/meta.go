package model

type HookConfigDefine struct {
	ContentType string `json:"content_type"`
	InsecureSsl string `json:"insecure_ssl"`
	Url         string `json:"url"`
}

type HookDefine struct {
	Name   string           `json:"name"`
	Active bool             `json:"active"`
	Events []string         `json:"events"`
	Config HookConfigDefine `json:"config"`
}

type Meta struct {
	Common
	Hook HookDefine `json:"hook"`
}
