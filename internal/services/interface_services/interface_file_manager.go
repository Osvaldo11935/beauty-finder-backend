package service_interface

import "github.com/gin-gonic/gin"

type IFileManager interface {
	Upload(ctx *gin.Context)(*string, error)
	Download(fileId string) (*string, error)
}