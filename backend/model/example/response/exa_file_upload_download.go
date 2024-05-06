package response

import "github.com/firwoodlin/letstrans/model/example"

type ExaFileResponse struct {
	File example.ExaFileUploadAndDownload `json:"file"`
}
