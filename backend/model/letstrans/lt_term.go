package letstrans

// Term .
type Term struct {
	BaseModel
	GlossaryID uint   `json:"glossary_id"` //术语库 ID
	SourceLang string `json:"source_lang"`
	TargetLang string `json:"target_lang"`
	SourceText string `json:"source_text"`
	TargetText string `json:"target_text"`
}
