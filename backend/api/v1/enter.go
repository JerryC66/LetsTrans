package v1

import (
	"github.com/firwoodlin/letstrans/api/v1/letstrans"
	"github.com/firwoodlin/letstrans/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup    system.ApiGroup
	LetsTransApiGroup letstrans.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
