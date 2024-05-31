package letstrans

// Document .
type Document struct {
	BaseModel
	Progress   float32 `json:"progress"`
	Author     string  `json:"author"`
	Filetype   string  `json:"filetype"`
	Name       string  `json:"name"`
	ProjectID  uint    `json:"project_id"`
	SourceLang string  `json:"source_lang"`
	TargetLang string  `json:"target_lang"`
	FilePath   string  `json:"-"`
	FileID     uint    `json:"file_id"`
	AuthorID   uint    `json:"author_id"`
}
