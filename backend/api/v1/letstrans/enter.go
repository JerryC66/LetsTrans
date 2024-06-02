package letstrans

import (
	"github.com/firwoodlin/letstrans/service/letstrans"
	"github.com/firwoodlin/letstrans/service/system"
)

type ApiGroup struct {
	FileOpApi
	ProjectApi
	SegmentApi
	GlossaryApi
	ThirdPartyApi
	TranslationMemoryApi
}

var (
	fileService        letstrans.FileService
	projectService     letstrans.ProjectService
	documentService    letstrans.DocumentService
	segmentService     letstrans.SegmentService
	glossaryService    letstrans.GlossaryService
	userService        system.UserService
	thirdPartyService  letstrans.ThirdPartyService
	transMemoryService letstrans.TranslateMemoryService
)
