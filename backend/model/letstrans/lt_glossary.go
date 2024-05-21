package letstrans

// Glossary
type Glossary struct {
	BaseModel
	// 作者昵称
	Author  string `json:"author"`
	Comment string `json:"comment"`
	// 术语库名称
	Name string `json:"name"`
}
