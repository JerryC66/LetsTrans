package letstrans

import "github.com/firwoodlin/letstrans/service/letstrans"

type ApiGroup struct {
	FileUploadApi
	ProjectApi
	SegmentApi
}

var (
	fileService     letstrans.FileService
	projectService  letstrans.ProjectService
	documentService letstrans.DocumentService
	segmentService  letstrans.SegmentService
)
