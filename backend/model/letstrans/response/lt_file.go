package response

import (
	"github.com/firwoodlin/letstrans/model/letstrans"
)

type ExaFileResponse struct {
	File letstrans.FileRecord `json:"file"`
}
