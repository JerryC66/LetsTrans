package letstrans

import (
	v1 "github.com/firwoodlin/letstrans/api/v1"
	"github.com/gin-gonic/gin"
)

type SegmentRouter struct{}

func (s *SegmentRouter) InitSegmentRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	segmentRouter := Router.Group("")
	segmentApi := v1.ApiGroupApp.LetsTransApiGroup.SegmentApi
	{
		segmentRouter.GET("/documents/:document_id/segments", segmentApi.GetSegmentList)
		segmentRouter.PUT("/documents/:document_id/segments/:segment_id", segmentApi.UpdateSegment)
	}
	return segmentRouter
}
