package model

type WorkflowDispatch struct {
	Common
	Inputs struct {
		Name string `json:"name"`
	} `json:"inputs"`
	Workflow string `json:"workflow"`
}
