package letstrans

import (
	"github.com/firwoodlin/letstrans/global"
	"github.com/firwoodlin/letstrans/model/common/response"
	"github.com/firwoodlin/letstrans/model/letstrans/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ThirdPartyApi struct {
}

func (t *ThirdPartyApi) TranslateDocument(c *gin.Context) {
	req := request.TranslateRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		//global.GVA_LOG.Error("TranslateDocument failed, when binding request", zap.Error(err), zap.Any("request", c.Request.Body))
		return
	}
	segs, err := thirdPartyService.MachineTranslate(req.DocumentID, req.Engine)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		global.GVA_LOG.Error("MachineTranslate failed, when calling thirdPartyService.MachineTranslate", zap.Error(err))
		return
	}
	response.OkWithData(gin.H{"segments": segs}, c)
}
