package app

import "template/config"

type (
	App struct {
		Config config.Config
	}

	// Config struct {
	// 	Port  string
	// 	Debug bool
	// }
)

func NewApp(cfg config.Config) *App {

	return &App{Config: cfg}
}

func (a *App) Init() error {

	// Here we initialize our services and any deps
	return nil
}
