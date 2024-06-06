package letstrans

import (
	"fmt"
	"github.com/firwoodlin/letstrans/global"
	"github.com/firwoodlin/letstrans/model/letstrans"
	"github.com/firwoodlin/letstrans/utils/upload"
	"go.uber.org/zap"
	"mime/multipart"
	"strings"
)

// FileService 定义文件服务结构体
type FileService struct{}

func (fileService *FileService) UploadFile(header *multipart.FileHeader) (file letstrans.FileRecord, err error) {
	oss := upload.NewOss()
	filePath, key, uploadErr := oss.UploadFile(header)
	if uploadErr != nil {
		//panic(uploadErr)
		global.GVA_LOG.Info("OSS 文件上传失败!", zap.Any("err", uploadErr))
		return
	}
	s := strings.Split(header.Filename, ".")
	fmt.Println(key, s)
	f := letstrans.FileRecord{
		FileName: header.Filename,
		FilePath: filePath,
		FileType: s[len(s)-1],
	}
	err = global.GVA_DB.Create(&f).Error
	if err != nil {
		global.GVA_LOG.Error("文件记录创建失败!", zap.Any("err", err))
		return f, err
	}
	return f, nil
}
