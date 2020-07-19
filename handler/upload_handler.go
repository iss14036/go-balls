package handler

import (
	"fmt"
	"github.com/tus/tusd/pkg/filestore"
	tusd "github.com/tus/tusd/pkg/handler"
	"go-balls/config"
	"log"
	"net/http"
)

type (
	UploadHandler struct {
		config *config.Config
	}
)

func NewUploadHandler(config *config.Config) *UploadHandler {
	return &UploadHandler{config: config}
}

func (uh *UploadHandler) Upload() http.Handler {
	store := filestore.FileStore{
		Path: "./assets/uploads",
	}

	composer := tusd.NewStoreComposer()
	store.UseIn(composer)

	handler, err := tusd.NewHandler(tusd.Config{
		BasePath:              "/files/",
		StoreComposer:         composer,
		NotifyCompleteUploads: true,
	})
	if err != nil {
		panic(fmt.Errorf("Unable to create handler: %s", err))
	}

	go func() {
		for {
			event := <-handler.CompleteUploads
			log.Println("upload...")
			fmt.Printf("Upload %s finished\n", event.Upload.ID)
		}
	}()

	return http.StripPrefix("/files/", handler)
}
