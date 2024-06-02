package letstrans

import (
	v1 "github.com/firwoodlin/letstrans/api/v1"
	"github.com/gin-gonic/gin"
)

type TranslationMemoryRouter struct{}

func (t *TranslationMemoryRouter) InitTranslationMemoryRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	translationMemoryRouter := Router.Group("")
	translationMemoryApi := v1.ApiGroupApp.LetsTransApiGroup.TranslationMemoryApi
	{
		translationMemoryRouter.GET("/memories/suggestion", translationMemoryApi.GetMemBySrc)

	}
	return translationMemoryRouter
}
