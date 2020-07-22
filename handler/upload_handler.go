package handler

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	tusd "github.com/tus/tusd/pkg/handler"
	"github.com/tus/tusd/pkg/s3store"
	"go-balls/common"
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
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		metadata := req.Header.Get("Upload-Metadata")
		if metadata != "" {
			fileType, err := common.GetFileType(metadata)
			if err != nil {
				log.Panic(err)
			}
			log.Println(fileType)

			if !common.IsFileTypeVideo(fileType) {
				log.Println("Error: Filetype is not a video")
				err := errors.New("Filetype is not a video")
				common.SendError(rw, err, http.StatusBadRequest)
				return
			}
		}

		sess := session.Must(session.NewSession(aws.NewConfig().
			WithMaxRetries(3),
		))

		// Create S3 service client with a specific Region.
		svc := s3.New(sess, aws.NewConfig().
			WithRegion(uh.config.AwsRegion),
		)

		store := s3store.New(uh.config.AwsBucket, svc)

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

		handler.ServeHTTP(rw, req)
	})
}
