package cli

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go-balls/appcontext"
	"go-balls/router"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

type (
	Cli struct {
		Args []string
	}
)

func NewCli(args []string) *Cli {
	return &Cli{Args: args}
}

func (c *Cli) Run(app *appcontext.Application) {
	log.SetLevel(log.InfoLevel)
	log.StandardLogger()
	log.SetOutput(os.Stdout)
	if strings.ToLower(app.Config.LogLevel) == "debug" {
		log.SetLevel(log.DebugLevel)
	}
	log.SetReportCaller(true)

	srv := 	&http.Server{
		Addr:	fmt.Sprintf(":%v", app.Config.AppPort),
		Handler: router.NewRouter(app),
	}

	log.Println(fmt.Sprintf("starting application { %v } on port :%v", app.Config.AppName, app.Config.AppPort))
	go listenServer(srv)
	waitForShutdown(srv)
}

func listenServer(apiServer *http.Server) {
	err := apiServer.ListenAndServe()
	if err != http.ErrServerClosed {
		log.WithField("error", err.Error()).Fatal("Server exited because of an error")
	}
}

func waitForShutdown(apiServer *http.Server) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig,
		syscall.SIGINT,
		syscall.SIGTERM)
	_ = <-sig
	log.Warn("Api server shutting down")
	if err := apiServer.Shutdown(context.Background()); err != nil {
		log.Println(err)
	}
	log.Warn("API server shutdown complete")
}