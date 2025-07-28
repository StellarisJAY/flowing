package main

import (
	"flag"
	"flowing/internal/config"
	"flowing/internal/repository"
	"gopkg.in/yaml.v3"
	"os"
)

func initialize(conf *config.Config) {
	// 初始化数据库
	repository.Init(conf)
}

func main() {
	confPath := flag.String("config", "config/config-dev.yml", "config file path")
	confFile, err := os.Open(*confPath)
	if err != nil {
		panic(err)
	}
	conf := config.Config{}
	if err := yaml.NewDecoder(confFile).Decode(&conf); err != nil {
		panic(err)
	}
	initialize(&conf)
}
