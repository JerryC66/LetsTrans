package letstrans

import (
	"github.com/firwoodlin/letstrans/global"
	"github.com/firwoodlin/letstrans/model/letstrans"
)

// GlossaryService 定义术语库服务结构体
type GlossaryService struct{}

func (s *GlossaryService) CreateGlossary(glossary letstrans.Glossary) (err error) {
	err = global.GVA_DB.Create(&glossary).Error
	return err
}

func (s *GlossaryService) GetGlossaryByID(glossaryID uint) (glossary letstrans.Glossary, err error) {
	err = global.GVA_DB.Model(&letstrans.Glossary{}).Where("id = ?", glossaryID).First(&glossary).Error
	return glossary, err
}

func (s *GlossaryService) GetGlossaryByUserID(uId uint) (glossaries []letstrans.Glossary, err error) {
	err = global.GVA_DB.Model(&letstrans.Glossary{}).Where("author_id = ?", uId).Find(&glossaries).Error
	return glossaries, err
}

func (s *GlossaryService) UpdateGlossary(glossaryID uint, glossary letstrans.Glossary) (err error) {
	err = global.GVA_DB.Model(&letstrans.Glossary{}).Where("id = ?", glossaryID).Updates(&glossary).Error
	return err
}

func (s *GlossaryService) DeleteGlossary(glossaryID uint) (err error) {
	err = global.GVA_DB.Delete(&letstrans.Glossary{}, glossaryID).Error
	return err
}

func (s *GlossaryService) CreateTerm(term letstrans.Term) (err error) {
	err = global.GVA_DB.Create(&term).Error
	return err
}

func (s *GlossaryService) GetTermsByGlossaryID(glossaryID uint) (terms []letstrans.Term, err error) {
	err = global.GVA_DB.Model(&letstrans.Term{}).Where("glossary_id = ?", glossaryID).Find(&terms).Error
	return terms, err
}

func (s *GlossaryService) UpdateTerm(termID uint, term letstrans.Term) (err error) {
	err = global.GVA_DB.Model(&letstrans.Term{}).Where("id = ?", termID).Updates(map[string]interface{}{
		"source_lang": term.SourceLang,
		"target_lang": term.TargetLang,
		"source_text": term.SourceText,
		"target_text": term.TargetText,
	}).Error
	return err
}

func (s *GlossaryService) DeleteTerm(termID uint) (err error) {
	err = global.GVA_DB.Delete(&letstrans.Term{}, termID).Error
	return err
}

func (s *GlossaryService) CreateTermInBatch(terms []letstrans.Term) (err error) {
	err = global.GVA_DB.Create(&terms).Error
	return err
}
