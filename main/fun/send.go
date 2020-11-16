package fun

import (
	"fmt"
	"strings"
)

type Message struct {
	event   string
	action  string
	repo    string
	repoUrl string
	user    string
	userUrl string
	detail  string
}

func SendMessage(message Message) {
	for _, value := range ParsedConfig.IM {
		if false == value.Enable {
			continue
		}
		content := value.TemplateContent
		content = strings.ReplaceAll(content, "${event}", message.event)
		content = strings.ReplaceAll(content, "${action}", message.action)
		content = strings.ReplaceAll(content, "${repo}", message.repo)
		content = strings.ReplaceAll(content, "${repoUrl}", message.repoUrl)
		content = strings.ReplaceAll(content, "${user}", message.user)
		content = strings.ReplaceAll(content, "${userUrl}", message.userUrl)
		content = strings.ReplaceAll(content, "${detail}", message.detail)
		fmt.Println(content)
	}
}
