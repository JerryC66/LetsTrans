package letstrans

// TranslationMemory
type TranslationMemory struct {
	//BaseModel
	DocumentID uint   `json:"document_id"`
	SourceText string `json:"source_text"`
	TargetText string `json:"target_text"`
	SimRank    int64  `json:"sim_rank"`
}
