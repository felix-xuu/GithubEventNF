package model

type SecurityAdvisoryDefine struct {
	Summary     string `json:"summary"`
	Description string `json:"description"`
	Severity    string `json:"severity"`
}

type SecurityAdvisory struct {
	Common
	SecurityAdvisory SecurityAdvisoryDefine `json:"security_advisory"`
}
