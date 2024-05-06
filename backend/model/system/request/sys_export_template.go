package request

import (
	"github.com/firwoodlin/letstrans/model/common/request"
	"github.com/firwoodlin/letstrans/model/system"
	"time"
)

type SysExportTemplateSearch struct {
	system.SysExportTemplate
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}
