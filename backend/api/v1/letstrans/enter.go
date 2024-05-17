package letstrans

import "github.com/firwoodlin/letstrans/service/letstrans"

type ApiGroup struct {
	FileUploadApi
}

var (
	fileService letstrans.FileService
)
