package letstrans

import "github.com/firwoodlin/letstrans/service/letstrans"

type ApiGroup struct {
	FileUploadApi
	ProjectApi
}

var (
	fileService     letstrans.FileService
	projectService  letstrans.ProjectService
	documentService letstrans.DocumentService
)
