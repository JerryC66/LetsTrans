package response

import "github.com/firwoodlin/letstrans/model/letstrans"

type GlossaryResponse struct {
	Glossary letstrans.Glossary `json:"glossary"`
	Terms    []letstrans.Term   `json:"terms"`
}
