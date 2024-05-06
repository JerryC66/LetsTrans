package request

import (
	"github.com/firwoodlin/letstrans/model/common/request"
	"github.com/firwoodlin/letstrans/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
