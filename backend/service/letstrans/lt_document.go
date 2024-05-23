package letstrans

import (
	"github.com/firwoodlin/letstrans/global"
	"github.com/firwoodlin/letstrans/model/letstrans"
)

type DocumentService struct{}

func (ds *DocumentService) GetDocumentsByProjID(projID uint) (docs []letstrans.Document, err error) {
	err = global.GVA_DB.Model(&letstrans.Document{}).Where("project_id = ?", projID).Find(&docs).Error
	return docs, err
}
