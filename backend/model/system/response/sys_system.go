package response

import "github.com/firwoodlin/letstrans/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
