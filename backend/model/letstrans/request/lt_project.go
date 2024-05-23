package request

import "time"

type ProjectCreateReq struct {
	Comment    string    `json:"comment"`
	Deadline   time.Time `json:"deadline"`
	Name       string    `json:"name"`
	SourceLang string    `json:"source_lang"`
	TargetLang string    `json:"target_lang"`
}
