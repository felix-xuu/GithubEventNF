package model

type PackageVersionDefine struct {
	Version  string `json:"version"`
	Summary  string `json:"summary"`
	Body     string `json:"body"`
	BodyHtml string `json:"body_html"`
	HtmlUrl  string `json:"html_url"`
}

type PackageDefine struct {
	Name           string               `json:"name"`
	PackageType    string               `json:"package_type"`
	HtmlUrl        string               `json:"html_url"`
	Owner          Owner                `json:"owner"`
	PackageVersion PackageVersionDefine `json:"package_version"`
}

type Package struct {
	Common
	Package PackageDefine `json:"package"`
}
