package letstrans

import (
	"errors"
	"github.com/firwoodlin/letstrans/global"
	"github.com/firwoodlin/letstrans/model/letstrans"
	"github.com/firwoodlin/letstrans/model/system"
	"github.com/firwoodlin/letstrans/utils"
	"github.com/firwoodlin/letstrans/utils/document"
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

func (p *ProjectService) ExportDocument(fileIDs []uint, eType string) (zipPath string, err error) {
	if eType != letstrans.ET_Original && eType != letstrans.ET_Translated {
		return "", errors.New("export type error")
	}
	// 获取文件路径
	var fileRecords []letstrans.FileRecord
	err = global.GVA_DB.Model(&letstrans.FileRecord{}).Where("id in (?)", fileIDs).Find(&fileRecords).Error
	if err != nil {
		return "", err
	}
	// 导出的文件名
	ext := ".zip"
	name := "export"
	zipName := name + "_" + time.Now().Format("20060102150405") + ext
	zipPath = "uploads/zip/" + zipName
	// 导出文件路径
	files := make([]string, len(fileRecords))
	for i, file := range fileRecords {
		files[i] = file.FilePath
	}
	err = utils.ZipFiles(zipPath, files, ".", ".")
	if err != nil {
		return "", err
	}
	return zipPath, nil
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
