package router

import (
	"github.com/firwoodlin/letstrans/router/letstrans"
	"github.com/firwoodlin/letstrans/router/system"
)

type RouterGroup struct {
	System    system.RouterGroup
	LetsTrans letstrans.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
