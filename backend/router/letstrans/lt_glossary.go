package letstrans

import (
	v1 "github.com/firwoodlin/letstrans/api/v1"
	"github.com/gin-gonic/gin"
)

type GlossaryRouter struct{}

func (g *GlossaryRouter) InitGlossaryRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	glossaryRouter := Router.Group("/glossaries")
	glossaryApi := v1.ApiGroupApp.LetsTransApiGroup.GlossaryApi
	{
		glossaryRouter.POST("", glossaryApi.CreateGlossary)
		glossaryRouter.GET("", glossaryApi.GetGlossaryList)

		glossaryRouter.GET("/:glossary_id", glossaryApi.GetGlossary)
		glossaryRouter.PUT("/:glossary_id", glossaryApi.UpdateGlossary)
		glossaryRouter.DELETE("/:glossary_id", glossaryApi.DeleteGlossary)

		glossaryRouter.POST("/:glossary_id/terms", glossaryApi.CreateTerm)
		glossaryRouter.GET("/:glossary_id/terms", glossaryApi.GetTermsByGlossary)
		glossaryRouter.PUT("/terms/:term_id", glossaryApi.UpdateTerm)
		glossaryRouter.DELETE("/:glossary_id/terms/:term_id", glossaryApi.DeleteTerm)
		glossaryRouter.POST("/:glossary_id/terms/batch", glossaryApi.CreateTermInBatch)
	}
	return glossaryRouter
}
