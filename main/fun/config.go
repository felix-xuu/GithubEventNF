package fun

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"path/filepath"
)

type Config struct {
	Server struct {
		Path string `yaml:"path"`
		Port string `yaml:"port"`
	} `yaml:"server"`
	IM map[string]struct {
		Enable          bool   `yaml:"enable"`
		TemplateFile    string `yaml:"template_file"`
		Token           string `yaml:"token"`
		ChatId          int64  `yaml:"chat_id"`
		MessageType     string `yaml:"message_type"`
		TemplateContent string
	} `yaml:"im"`
}

var ParsedConfig Config

func ConfigParse() {
	file, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Panic(err)
	}
	err = yaml.Unmarshal(file, &ParsedConfig)
	if err != nil {
		log.Panic(err)
	}
	if "" == ParsedConfig.Server.Path {
		ParsedConfig.Server.Path = "/hook"
	}
	if "" == ParsedConfig.Server.Port {
		ParsedConfig.Server.Port = "8080"
	}
	var enabledIM []string
	for key, value := range ParsedConfig.IM {
		if true == value.Enable {
			if "" == value.TemplateFile {
				value.TemplateFile = "default.txt"
			}
			enabledIM = append(enabledIM, key)
			file, err := ioutil.ReadFile(filepath.Join("templates", value.TemplateFile))
			if err != nil {
				log.Panic(err)
			}
			value.TemplateContent = string(file)
			ParsedConfig.IM[key] = value
		}
	}
	log.Println("enabled IMs:", enabledIM)
}
