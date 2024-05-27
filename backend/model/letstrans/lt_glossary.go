package letstrans

// Glossary
type Glossary struct {
	BaseModel
	Name       string `json:"name"`   // 术语库名称
	Author     string `json:"author"` // 作者昵称
	Comment    string `json:"comment"`
	SourceLang string `json:"source_lang"`
	TargetLang string `json:"target_lang"`
	AuthorID   uint   `json:"author_id"`
}
