package appcontext

import "go-balls/config"

type Application struct {
	Config config.Config
}

func NewApp() *Application {
	return &Application{
		Config: config.AppConfig(),
	}
}

