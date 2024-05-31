package letstrans

import (
	v1 "github.com/firwoodlin/letstrans/api/v1"
	"github.com/gin-gonic/gin"
)

type ThirdPartyRouter struct{}

func (t *ThirdPartyRouter) InitThirdPartyRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	thirdPartyRouter := Router.Group("")
	thirdPartyApi := v1.ApiGroupApp.LetsTransApiGroup.ThirdPartyApi
	{
		thirdPartyRouter.POST("/machine/translate", thirdPartyApi.TranslateDocument)
	}
	return thirdPartyRouter
}
