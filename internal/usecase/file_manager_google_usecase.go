package usecase

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/url"
	"os"
	"src/internal/configs"
	"src/internal/domain/errors"
	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"google.golang.org/appengine"
)          

func Upload(context *gin.Context) (*string, error) {

	config, loadConfigErr := configs.LoadConfig()

	if loadConfigErr != nil {
		return nil, loadConfigErr
	}
	if _, err := os.Stat(config.FileKeys); os.IsNotExist(err) {
		return nil, errors.ConfigurationFileNotFoundError()
	}
	ctx := appengine.NewContext(context.Request)

	storageClient, err := storage.NewClient(ctx, option.WithCredentialsFile(config.FileKeys))

	if err != nil {
		return nil, errors.NewGoogleStorageUnknownError(err.Error())
	}

	f, uploadedFile, err := context.Request.FormFile("file")

	if err != nil {
		return nil, errors.NewGoogleStorageUnknownError(err.Error())
	}

	defer func(f multipart.File) {
		if err := f.Close(); err != nil {
			err := errors.NewGoogleStorageUnknownError(err.Error())
			if err != nil {
				return
			}
		}
	}(f)

	sw := storageClient.Bucket(config.Bucket).Object(uploadedFile.Filename).NewWriter(ctx)

	if _, err := io.Copy(sw, f); err != nil {
		return nil, errors.NewGoogleStorageUnknownError(err.Error())
	}

	if err := sw.Close(); err != nil {
		return nil, errors.NewGoogleStorageUnknownError(err.Error())
	}

	u, err := url.Parse(fmt.Sprintf("%s/%s/%s", config.BaseUrlGoogleStorage, config.Bucket, uploadedFile.Filename))

	if err != nil {
		return nil, errors.NewGoogleStorageUnknownError(err.Error())
	}

	linkFile := u.String()

	return &linkFile, nil
}
