package letstrans

import (
	"github.com/firwoodlin/letstrans/global"
	"github.com/firwoodlin/letstrans/model/common/response"
	"github.com/firwoodlin/letstrans/model/letstrans"
	letRes "github.com/firwoodlin/letstrans/model/letstrans/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type FileUploadApi struct{}

func (f *FileUploadApi) UploadFile(c *gin.Context) {
	var file letstrans.FileRecord
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		global.GVA_LOG.Error("接收文件失败!", zap.Error(err))
		response.FailWithMessage("接收文件失败", c)
		return
	}
	file, err = fileService.UploadFile(header) // 文件上传后拿到文件路径
	if err != nil {
		global.GVA_LOG.Error("修改数据库链接失败!", zap.Error(err))
		response.FailWithMessage("修改数据库链接失败", c)
		return
	}
	response.OkWithDetailed(letRes.ExaFileResponse{File: file}, "上传成功", c)
}
