package request

import (
	"github.com/firwoodlin/letstrans/model/common/request"
	"github.com/firwoodlin/letstrans/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
