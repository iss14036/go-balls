package config

import (
	"fmt"
	"log"
	"os"

	"github.com/subosito/gotenv"
)

type (
	Config struct {
		AppName     string
		AppPort     string
		LogLevel    string
		Environment string
		AwsBucket   string
		AwsRegion   string
	}
)

var appConfig = Config{}

func init() {
	if err := gotenv.Load(); err != nil {
		log.Println("Loading from os env variable")
	}
	log.SetOutput(os.Stdout)
	appConfig = NewConfig()
}

func NewConfig() Config {
	c := Config{
		AppName:     GetString("APP_NAME"),
		AppPort:     GetString("APP_PORT"),
		LogLevel:    GetString("LOG_LEVEL"),
		Environment: GetString("ENVIRONMENT"),
		AwsBucket:   GetString("AWS_BUCKET"),
		AwsRegion:   GetString("AWS_REGION"),
	}
	log.Println(fmt.Sprintf("configuration loaded\n %#v", c))
	return c
}

func AppConfig() Config {
	return appConfig
}
