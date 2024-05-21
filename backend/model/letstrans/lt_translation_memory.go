package letstrans

// TranslationMemory
type TranslationMemory struct {
	BaseModel
	DocumentID string `json:"document_id"`
	SourceText string `json:"source_text"`
	TargetText string `json:"target_text"`
}
