package letstrans

// Term .
type Term struct {
	BaseModel
	DocumentID string `json:"document_id"`
	SourceLang string `json:"source_lang"`
	SourceText string `json:"source_text"`
	TargetLang string `json:"target_lang"`
	TargetText string `json:"target_text"`
}
