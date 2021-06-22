package config

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

type Server struct {
	Port int `yaml:"port"`
}

type Gitee struct {
	TokenUrl     string `yaml:"token-url"`
	ClientId     string `yaml:"client-id"`
	RedirectUri  string `yaml:"redirect-uri"`
	ClientSecret string `yaml:"client-secret"`
}

type Github struct {
	TokenUrl     string `yaml:"token-url"`
	ClientId     string `yaml:"client-id"`
	RedirectUri  string `yaml:"redirect-uri"`
	ClientSecret string `yaml:"client-secret"`
}

type Web struct {
	IndexUrl string `yaml:"index-url"`
	ErrorUrl string `yaml:"web-url"`
}

type Config struct {
	ServerConfig Gitee  `yaml:"server"`
	GiteeConfig  Gitee  `yaml:"gitee"`
	GithubConfig Github `yaml:"github"`
	WebConfig    Web    `yaml:"web"`
}

var (
	AppConfig Config
)

func InitAppConfig() {
	file, _ := ioutil.ReadFile("./config.yml")
	if err := yaml.Unmarshal(file, &AppConfig); err != nil {
		log.Println("ERROR", err)
	}
}
