package letstrans

import (
	v1 "github.com/firwoodlin/letstrans/api/v1"
	"github.com/gin-gonic/gin"
)

type FileRouter struct{}

func (f *FileRouter) InitFileRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	fileRouter := Router.Group("files")
	fileApi := v1.ApiGroupApp.LetsTransApiGroup.FileUploadApi
	{
		fileRouter.POST("", fileApi.UploadFile)
	}
	return fileRouter
}
