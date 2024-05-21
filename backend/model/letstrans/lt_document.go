package letstrans

// Document .
type Document struct {
	BaseModel
	// 创建者
	Author string `json:"author"`
	// 文件类型
	Filetype string `json:"filetype"`
	// UUID
	Name string `json:"name"`
	// 翻译进度，0-1 之间的数字
	Progress   float64 `json:"progress"`
	SourceLang string  `json:"source_lang"`
	TargetLang string  `json:"target_lang"`
}
