package model

type PusherDefine struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CommitDefine struct {
	Id        string `json:"id"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
	Url       string `json:"url"`
}

type Push struct {
	Common
	Ref     string         `json:"ref"`
	Pusher  PusherDefine   `json:"pusher"`
	Compare string         `json:"compare"`
	Commits []CommitDefine `json:"commits"`
}
