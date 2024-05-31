package letstrans

import (
	"github.com/firwoodlin/letstrans/model/common/response"
	"github.com/firwoodlin/letstrans/model/letstrans/request"
	"github.com/gin-gonic/gin"
)

type ThirdPartyApi struct {
}

func (t *ThirdPartyApi) TranslateDocument(c *gin.Context) {
	req := request.TranslateRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	segs, err := thirdPartyService.MachineTranslate(req.DocumentID, req.Engine)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(gin.H{"segments": segs}, c)
}
