package fun

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
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
	for key, value := range ParsedConfig.IM {
		if false == value.Enable {
			continue
		}
		content := value.TemplateContent
		content = strings.ReplaceAll(content, "{{event}}", message.event)
		content = strings.ReplaceAll(content, "{{action}}", message.action)
		content = strings.ReplaceAll(content, "{{repo}}", message.repo)
		content = strings.ReplaceAll(content, "{{repoUrl}}", message.repoUrl)
		content = strings.ReplaceAll(content, "{{user}}", message.user)
		content = strings.ReplaceAll(content, "{{userUrl}}", message.userUrl)
		content = strings.ReplaceAll(content, "{{detail}}", message.detail)
		log.Println(content)
		switch key {
		case Telegram:
			sendToTelegram(content)
		case DingTalk:
			sendToDingTalk(content)
		default:
		}
	}
}

func sendToTelegram(message string) {
	bot, err := tgbotapi.NewBotAPI(ParsedConfig.IM[Telegram].Token)
	if err != nil {
		log.Panic(err)
	}
	msg := tgbotapi.NewMessage(ParsedConfig.IM[Telegram].ChatId, message)
	msg.ParseMode = ParsedConfig.IM[Telegram].MessageType
	_, err = bot.Send(msg)
	if err != nil {
		log.Panic(err)
	}
}

func sendToDingTalk(message string) {

}
