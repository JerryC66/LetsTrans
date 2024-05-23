package response

import "github.com/firwoodlin/letstrans/model/letstrans"

type ProjectDetailRes struct {
	Documents []letstrans.Document `json:"documents"`
	Project   letstrans.Project    `json:"project"`
}
