package handler

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	tusd "github.com/tus/tusd/pkg/handler"
	"github.com/tus/tusd/pkg/s3store"
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
