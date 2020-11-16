package model

type Ping struct {
	Common
	Zen  string     `json:"zen"`
	Hook HookDefine `json:"hook"`
}
