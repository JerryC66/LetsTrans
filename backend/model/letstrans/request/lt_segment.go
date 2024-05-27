package request

type SegmentUpdateRequest struct {
	Finished   *bool  `json:"finished"`
	TargetText string `json:"target_text"`
}
