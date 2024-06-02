package letstrans

import (
	"fmt"
	"github.com/firwoodlin/letstrans/global"
	"github.com/firwoodlin/letstrans/model/letstrans"
	"github.com/firwoodlin/letstrans/model/system"
	"github.com/firwoodlin/letstrans/utils"
	"github.com/firwoodlin/letstrans/utils/document"
	"go.uber.org/zap"
	"time"
)

// ProjectService 定义项目服务结构体
type ProjectService struct{}

func (p *ProjectService) CreateProject(req letstrans.Project) (proj letstrans.Project, err error) {
	err = global.GVA_DB.Model(&letstrans.Project{}).Create(&req).Error
	return req, nil
}

func (p *ProjectService) GetProjectList(userID uint) (projects []letstrans.Project, err error) {
	err = global.GVA_DB.Model(&letstrans.Project{}).Where("user_id = ?", userID).Find(&projects).Error
	return projects, err
}

func (p *ProjectService) GetProjectByID(projID uint) (project letstrans.Project, err error) {
	err = global.GVA_DB.Model(&letstrans.Project{}).Where("id = ?", projID).First(&project).Error
	return project, err
}

// AddDocument 添加文档, 同时开启 doc to seg 任务
func (p *ProjectService) AddDocument(fileID uint, authorID uint, projID uint) (doc *letstrans.Document, err error) {
	file := letstrans.FileRecord{}
	err = global.GVA_DB.Model(&letstrans.FileRecord{}).Where("id = ?", fileID).First(&file).Error
	if err != nil {
		return nil, err
	}
	author := system.SysUser{}
	err = global.GVA_DB.Model(&system.SysUser{}).Where("id = ?", authorID).First(&author).Error
	if err != nil {
		return nil, err
	}
	proj := letstrans.Project{}
	err = global.GVA_DB.Model(&letstrans.Project{}).Where("id = ?", projID).First(&proj).Error
	if err != nil {
		return nil, err
	}
	doc = &letstrans.Document{
		Name:       file.FileName,
		Author:     author.NickName, // 昵称作为作者
		AuthorID:   authorID,
		Filetype:   file.FileType,
		ProjectID:  projID,
		FilePath:   file.FilePath,
		FileID:     fileID,
		SourceLang: proj.SourceLang,
		TargetLang: proj.TargetLang,
	}
	err = global.GVA_DB.Model(&letstrans.Document{}).Create(&doc).Error
	if err != nil {
		return
	}
	// 开启 doc to seg 任务
	segments, err := document.ProcessDocument(*doc)
	if err != nil {
		return
	}
	for i := range segments {
		segments[i].DocumentID = doc.ID
	}
	err = global.GVA_DB.Model(&letstrans.Segment{}).Create(&segments).Error
	if err != nil {
		return
	}
	return
}

// DeleteDocuments 删除文档
func (p *ProjectService) DeleteDocuments(fileIDs []uint, projID uint) (err error) {
	err = global.GVA_DB.Model(&letstrans.Document{}).Where("file_id in (?) AND project_id = ?", fileIDs, projID).Delete(&letstrans.Document{}).Error
	return err
}

func (p *ProjectService) DeleteProject(projIDs []uint) (err error) {
	err = global.GVA_DB.Model(&letstrans.Project{}).Where("id in (?)", projIDs).Delete(&letstrans.Project{}).Error
	return err
}

func (p *ProjectService) ExportDocument(fileIDs []uint, exportType string) (zipPath string, err error) {
	// 获取文件路径
	var documents []letstrans.Document
	err = global.GVA_DB.Model(&letstrans.Document{}).Where("id in (?)", fileIDs).Find(&documents).Error
	if err != nil {
		return "", err
	}
	// 导出的文件名
	ext := ".zip"
	name := "export"
	zipName := name + "_" + exportType + "_" + time.Now().Format("20060102150405") + ext
	zipPath = "uploads/zip/" + zipName
	// 导出文件路径
	files := make([]string, len(documents))
	switch exportType {
	case letstrans.ET_Original:
		files = exportOriginalFile(documents)
	case letstrans.ET_Translated:
		files = exportTranslatedFile(documents)
	default:
		return "", fmt.Errorf("unsupported export type: %s", exportType)
	}
	err = utils.ZipFiles(zipPath, files, ".", ".")
	if err != nil {
		return "", err
	}
	return zipPath, nil
}

func exportOriginalFile(docs []letstrans.Document) (files []string) {
	files = make([]string, len(docs))
	for i, record := range docs {
		files[i] = record.FilePath
	}
	return files
}
func exportTranslatedFile(docs []letstrans.Document) (files []string) {
	//document.ModifyDocument
	files = make([]string, len(docs))
	for i, doc := range docs {
		// 获取文档的所有段落
		var segments []letstrans.Segment
		err := global.GVA_DB.Model(&letstrans.Segment{}).Where("document_id = ?", doc.ID).Find(&segments).Error
		if err != nil {
			global.GVA_LOG.Error("Get segments failed", zap.Error(err))
			continue
		}
		// 获取替换后的文档
		filePath, err := document.ModifyDocument(doc.FilePath, segments)
		if err != nil {
			global.GVA_LOG.Error("Modify document failed", zap.Error(err))
			continue
		}
		files[i] = filePath
	}
	return files
}

//func (p *ProjectService) GetProjectDetail(projID uint) (res response.ProjectDetailRes, err error) {
//	var project letstrans.Project
//	var docs []letstrans.Document
//	err = global.GVA_DB.Model(&letstrans.Project{}).
//}

//func (p *ProjectService) UpdateProject(proj letstrans.Project) (err error) {
//	err = global.GVA_DB.Model(&letstrans.Project{}).Where("id = ?", proj.ID).Updates(&proj).Error
//	return err
//}
