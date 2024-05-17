package letstrans

// Segment .
type Segment struct {
	BaseModel
	// 关联的文档 ID
	DocumentID string `json:"document_id"`
	// 是否已经标记完成
	Finished bool `json:"finished"`
	// 自增整数，保证有序性
	SourceText string `json:"source_text"`
	TargetText string `json:"target_text"`
}
