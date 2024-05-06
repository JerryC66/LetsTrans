package router

import (
	"github.com/firwoodlin/letstrans/router/system"
)

type RouterGroup struct {
	System system.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
