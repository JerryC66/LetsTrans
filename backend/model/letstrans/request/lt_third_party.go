package request

type TranslateRequest struct {
	DocumentID uint   `json:"document_id"`
	Engine     string `json:"engine"`
}
