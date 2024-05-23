package letstrans

import "time"

// Project .
type Project struct {
	BaseModel
	// 注释
	Comment string `json:"comment"`
	// 截止日期
	Deadline time.Time `json:"deadline"`
	// 项目名称
	Name string `json:"name"`
	// 翻译进度
	Progress   float64 `json:"progress"`
	SourceLang string  `json:"source_lang"`
	TargetLang string  `json:"target_lang"`
	UserID     uint    `json:"user_id"`
}
