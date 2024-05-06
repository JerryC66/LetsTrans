package response

import (
	"github.com/firwoodlin/letstrans/model/system/request"
)

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
