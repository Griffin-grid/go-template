package main

import (
	"context"

	"template/app"
	"template/config"
	"template/server"

	log "github.com/sirupsen/logrus"
)

func main() {

	cfg := config.NewConfig()

	err := cfg.LoadConfig()
	if err != nil {
		log.Errorf("an error occured while load config: %v", err)
	}

	a := app.NewApp(*cfg)

	err = a.Init()
	if err != nil {
		log.Errorf("an error occured while init app: %v", err)
	}

	router := a.InitRouter()

	srv := server.NewServer(a.Config.App.Port, router)

	err = srv.Start()
	if err != nil {
		log.Errorf("an error occured while start server: %v", err)
	}

	defer srv.Stop(context.TODO())
}
