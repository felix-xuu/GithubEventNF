package fun

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
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
		TemplateContent string
	} `yaml:"im"`
}

var ParsedConfig Config

func ConfigParse() {
	file, err := ioutil.ReadFile("config.yml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(file, &ParsedConfig)
	if err != nil {
		panic(err)
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
				panic(err)
			}
			value.TemplateContent = string(file)
			ParsedConfig.IM[key] = value
		}
	}
	fmt.Println("enabled IMs:", enabledIM)
}
