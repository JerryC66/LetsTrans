package letstrans

import (
	v1 "github.com/firwoodlin/letstrans/api/v1"
	"github.com/gin-gonic/gin"
)

type GlossaryRouter struct{}

func (g *GlossaryRouter) InitGlossaryRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	glossaryRouter := Router.Group("")
	glossaryApi := v1.ApiGroupApp.LetsTransApiGroup.GlossaryApi
	{
		glossaryRouter.POST("/glossaries", glossaryApi.CreateGlossary)
		glossaryRouter.GET("/glossaries", glossaryApi.GetGlossaryList)

		glossaryRouter.GET("/glossaries/:glossary_id", glossaryApi.GetGlossary)
		glossaryRouter.PUT("/glossaries/:glossary_id", glossaryApi.UpdateGlossary)
		glossaryRouter.DELETE("/glossaries/:glossary_id", glossaryApi.DeleteGlossary)

		glossaryRouter.POST("/glossaries/:glossary_id/terms", glossaryApi.CreateTerm)
		glossaryRouter.GET("/glossaries/:glossary_id/terms", glossaryApi.GetTermsByGlossary)
		glossaryRouter.PUT("/glossaries/terms/:term_id", glossaryApi.UpdateTerm)
		glossaryRouter.DELETE("/glossaries/:glossary_id/terms/:term_id", glossaryApi.DeleteTerm)
	}
	return glossaryRouter
}
