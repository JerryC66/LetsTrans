package letstrans

import (
	"github.com/firwoodlin/letstrans/global"
	comReq "github.com/firwoodlin/letstrans/model/common/request"
	"github.com/firwoodlin/letstrans/model/common/response"
	"github.com/firwoodlin/letstrans/model/letstrans"
	letsReq "github.com/firwoodlin/letstrans/model/letstrans/request"
	"github.com/firwoodlin/letstrans/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"path/filepath"
	"strconv"
)

type ProjectApi struct{}

func (pa *ProjectApi) CreateProject(c *gin.Context) {
	jwtId := utils.GetUserID(c)

	// 实现创建项目的逻辑
	var req letsReq.ProjectCreateReq
	// 从请求中解析出项目信息
	err := c.ShouldBindJSON(&req)
	if err != nil {
		// 如果解析失败, 则返回错误信息
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 调用服务层的创建项目方法
	// 传入项目信息, 返回创建的项目信息和错误信息
	proj, err := projectService.CreateProject(letstrans.Project{
		Name:       req.Name,
		UserID:     jwtId,
		Comment:    req.Comment,
		Deadline:   req.Deadline,
		SourceLang: req.SourceLang,
		TargetLang: req.TargetLang,
	})
	if err != nil {
		// 如果创建失败, 则返回错误信息
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 如果创建成功, 则返回成功信息和创建的项目信息
	response.OkWithData(proj, c)
}

func (pa *ProjectApi) GetProjects(c *gin.Context) {
	jwtId := utils.GetUserID(c)

	// 实现获取项目列表的逻辑
	// 调用服务层的获取项目列表方法
	// 传入用户ID, 返回项目列表和错误信息
	projects, err := projectService.GetProjectList(jwtId)
	if err != nil {
		// 如果获取失败, 则返回错误信息
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 如果获取成功, 则返回成功信息和项目列表
	response.OkWithData(projects, c)
}

// GetProjectByID 查询项目详情。
func (pa *ProjectApi) GetProjectByID(c *gin.Context) {
	//jwtId := utils.GetUserID(c)
	projIDStr := c.Param("project_id")
	// 从请求中解析出项目ID
	projID, err := strconv.ParseUint(projIDStr, 10, 32)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	proj, err := projectService.GetProjectByID(uint(projID))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	docs, err := documentService.GetDocumentsByProjID(uint(projID))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(gin.H{
		"project":   proj,
		"documents": docs,
	}, c)
}

func (pa *ProjectApi) DeleteProject(c *gin.Context) {
	//jwtId := utils.GetUserID(c)
	// 需要删除的 project id s
	var ids comReq.IdsReq
	err := c.ShouldBindJSON(&ids)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = projectService.DeleteProject(ids.Ids)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

func (pa *ProjectApi) AddDocument(c *gin.Context) {
	jwtId := utils.GetUserID(c)
	//authorInfo, err := userService.FindUserById(int(jwtId))
	//if err != nil {
	//	response.FailWithMessage(err.Error(), c)
	//	return
	//}

	projId := utils.Param2Uint(c, "project_id")
	if projId == 0 {
		response.FailWithMessage("project id error", c)
		return
	}

	// save file , get file ids , convert to document
	form, _ := c.MultipartForm()
	files := form.File["files"]
	fileRecords := make([]letstrans.FileRecord, len(files))
	docs := make([]*letstrans.Document, len(files))
	for i, file := range files {
		var err error = nil
		fileRecords[i], err = fileService.UploadFile(file) // 文件上传后拿到文件路径
		if err != nil {
			global.GVA_LOG.Error("修改数据库链接失败!", zap.Error(err))
			response.FailWithMessage("修改数据库链接失败", c)
			return
		}
		docs[i], err = projectService.AddDocument(fileRecords[i].ID, jwtId, projId)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
	}

	response.OkWithData(docs, c)
}

func (pa *ProjectApi) DeleteDocuments(c *gin.Context) {
	//jwtId := utils.GetUserID(c)
	projID := c.Param("project_id")
	// 从请求中解析出项目ID
	projIDUint, err := strconv.ParseUint(projID, 10, 32)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var ids comReq.IdsReq
	err = c.ShouldBindJSON(&ids)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = projectService.DeleteDocuments(ids.Ids, uint(projIDUint))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

func (pa *ProjectApi) DownloadDocuments(c *gin.Context) {
	req := letsReq.FileDownloadReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if req.Type != letstrans.ET_Original && req.Type != letstrans.ET_Translated {
		response.FailWithMessage("export type error", c)
		return
	}
	// 这里的 ID 实际上是 Document ID
	zipPath, err := projectService.ExportDocument(req.IDs, req.Type)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	c.Header("Content-Disposition", "attachment; filename="+filepath.Base(zipPath))

	c.File(zipPath)

}
