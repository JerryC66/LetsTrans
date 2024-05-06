package service

import "github.com/firwoodlin/letstrans/service/system"

type ServiceGroup struct {
	SystemServiceGroup system.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
