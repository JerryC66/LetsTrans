package letstrans

import (
	v1 "github.com/firwoodlin/letstrans/api/v1"
	"github.com/gin-gonic/gin"
)

type ProjectRouter struct{}

// CreateProject 创建项目
func (p *ProjectRouter) InitProjectRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	projectRouter := Router.Group("projects")
	projectApi := v1.ApiGroupApp.LetsTransApiGroup.ProjectApi
	{
		projectRouter.POST("", projectApi.CreateProject)
		projectRouter.GET("", projectApi.GetProjects)
		projectRouter.GET("/:project_id", projectApi.GetProjectByID)
		projectRouter.DELETE("", projectApi.DeleteProject)
		projectRouter.POST("/:project_id/files", projectApi.AddDocument)
		projectRouter.DELETE("/:project_id/files", projectApi.DeleteDocuments)
		projectRouter.GET("/:project_id/files", projectApi.DownloadDocuments)
	}
	return projectRouter
}
