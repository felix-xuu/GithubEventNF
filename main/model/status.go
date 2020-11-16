package model

type Status struct {
	Common
	Name        string `json:"name"`
	TargetUrl   string `json:"target_url"`
	Context     string `json:"context"`
	Description string `json:"description"`
	State       string `json:"state"`
}
