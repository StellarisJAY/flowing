package main

import (
	"flag"
	"flowing/api"
	"flowing/internal/config"
	"flowing/internal/migration"
	"flowing/internal/repository"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"net/http"
	"os"
	"os/signal"
)

func initialize(conf *config.Config) {
	// 初始化数据库
	repository.Init(conf)
	migration.MigrateDB()
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

	e := gin.New()
	api.InitRouter(e)

	go func() {
		if err := http.ListenAndServe(":"+conf.Server.Port, e); err != nil {
			panic(err)
		}
	}()
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, os.Kill)
	<-signalChan
}
