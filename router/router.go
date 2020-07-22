package router

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go-balls/appcontext"
	"go-balls/handler"
	"net/http"
	"os"
	"strings"
)

func NewRouter(app *appcontext.Application) http.Handler {
	uploadHandler := handler.NewUploadHandler(&app.Config)

	r := mux.NewRouter()
	tusUploadRouter(r, uploadHandler)
	publicRouter(r)

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)

	if strings.ToLower(app.Config.Environment) == "production" {
		return handlers.RecoveryHandler()(loggedRouter)
	}
	return loggedRouter
}

func tusUploadRouter(mainRouter *mux.Router, uh *handler.UploadHandler) {
	mainRouter.PathPrefix("/files/").Handler(http.StripPrefix("/files/", uh.Upload()))
}

func publicRouter(mainRouter *mux.Router) {
	mainRouter.HandleFunc("/", handler.HealthCheck).Methods("GET")
	mainRouter.HandleFunc("/ping", handler.HealthCheck).Methods("GET")
}