package letstrans

import (
	"github.com/firwoodlin/letstrans/model/common/response"
	"github.com/firwoodlin/letstrans/model/letstrans"
	"github.com/firwoodlin/letstrans/model/letstrans/request"
	"github.com/firwoodlin/letstrans/utils"
	"github.com/gin-gonic/gin"
)

type SegmentApi struct{}

// GetSegmentList GetSegmentList
func (a *SegmentApi) GetSegmentList(c *gin.Context) {
	docID := utils.Param2Uint(c, "document_id")
	if docID == 0 {
		response.FailWithMessage("docID is invalid", c)
		return
	}

	segments, err := segmentService.GetSegmentList(docID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	doc, err := documentService.GetDocumentByID(docID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(gin.H{
		"segments": segments,
		"doc":      doc,
	}, c)
}

// UpdateSegment UpdateSegment
func (a *SegmentApi) UpdateSegment(c *gin.Context) {
	var seg request.SegmentUpdateRequest
	err := c.ShouldBindJSON(&seg)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	segID := utils.Param2Uint(c, "segment_id")
	err = segmentService.UpdateSegment(segID, letstrans.Segment{Finished: seg.Finished, TargetText: seg.TargetText})
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}
