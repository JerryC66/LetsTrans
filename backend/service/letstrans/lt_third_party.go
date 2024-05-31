package letstrans

import (
	"github.com/firwoodlin/letstrans/global"
	"github.com/firwoodlin/letstrans/model/letstrans"
	"github.com/firwoodlin/letstrans/utils/translate"
)

// ThirdPartyService 定义第三方服务结构体
type ThirdPartyService struct{}

// MachineTranslate 批量进行机器翻译，将翻译结果存入数据库
func (tps *ThirdPartyService) MachineTranslate(documentID uint, engineType string) (segments []letstrans.Segment, err error) {
	// 实现批量进行机器翻译的逻辑
	var doc letstrans.Document
	err = global.GVA_DB.Model(&letstrans.Document{}).Where("id = ?", documentID).First(&doc).Error
	if err != nil {
		return
	}
	err = global.GVA_DB.Model(&letstrans.Segment{}).Where("document_id = ?", documentID).Order("id ASC").Find(&segments).Error
	if err != nil {
		return
	}
	engine := translate.NewEngine(engineType)
	for i, seg := range segments {
		// 此处调用翻译引擎进行翻译
		segments[i].TargetText, err = engine.Translate(seg.SourceText, doc.SourceLang, doc.TargetLang)
		err = global.GVA_DB.Model(&letstrans.Segment{}).Select("target_text").Where("id = ?", segments[i].ID).Updates(&segments[i]).Error
		if err != nil {
			return
		}
	}
	return
}
