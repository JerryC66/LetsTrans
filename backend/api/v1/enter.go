package v1

import (
	"github.com/firwoodlin/letstrans/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup system.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
