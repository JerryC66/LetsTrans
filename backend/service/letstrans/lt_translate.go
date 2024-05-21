package letstrans

import "time"

// TranslateBasisService 定义翻译基础服务结构体
type TranslateBasisService struct{}

// Segment 定义段落结构体
type Segment struct {
	DocumentID string    `json:"document_id"`
	ID         int       `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	SourceText string    `json:"source_text"`
	TargetText string    `json:"target_text"`
	Finished   bool      `json:"finished"`
}

// TMTrans 定义翻译记忆库结构体
type TMTrans struct {
	DocumentID string    `json:"document_id"`
	ID         string    `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	SourceText string    `json:"source_text"`
	TargetText string    `json:"target_text"`
	Percentage int       `json:"percentage"`
}

// Term 定义术语结构体
type Term struct {
	DocumentID string    `json:"document_id"`
	ID         string    `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	SourceLang string    `json:"source_lang"`
	TargetLang string    `json:"target_lang"`
	SourceText string    `json:"source_text"`
	TargetText string    `json:"target_text"`
}

// GetSegments 获取文档的所有segments
func (tbs *TranslateBasisService) GetSegments(documentID string) (Document, []Segment, error) {
	// 实现获取文档所有segments的逻辑
	return Document{}, []Segment{}, nil
}

// UpdateSegment 修改segment的属性，保存翻译进度
func (tbs *TranslateBasisService) UpdateSegment(documentID string, segmentID string, targetText string, finished bool) error {
	// 实现修改segment属性的逻辑
	return nil
}

// GetSegmentSuggestions 获取segment的记忆库和术语库
func (tbs *TranslateBasisService) GetSegmentSuggestions(documentID string, segmentID string) ([]TMTrans, []Term, error) {
	// 实现获取segment记忆库和术语库的逻辑
	return []TMTrans{}, []Term{}, nil
}
