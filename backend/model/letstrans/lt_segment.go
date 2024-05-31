package letstrans

// Segment .
type Segment struct {
	BaseModel
	DocumentID uint   `json:"document_id"`                   // 关联的文档 ID
	RelativeID uint   `json:"relative_id"`                   // 相对于文档的 ID
	Finished   *bool  `json:"finished" gorm:"default:false"` // 是否已经标记完成
	SourceText string `json:"source_text"`
	TargetText string `json:"target_text"`
}
