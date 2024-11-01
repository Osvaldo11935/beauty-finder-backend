package service_interface

import "github.com/gin-gonic/gin"

type IFileManager interface {
	Upload(ctx *gin.Context, filePath *string)(*string, error)
	GetFileUrl(fileId string) (*string, error)
	Download(filePath string) ([]byte, error)
}