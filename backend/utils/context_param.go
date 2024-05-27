package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func Param2Uint(c *gin.Context, key string) uint {
	valStr := c.Param(key)
	if valStr != "" {
		valUint, err := strconv.ParseInt(valStr, 10, 64)
		if err == nil {
			return uint(valUint)
		}
	}
	//else {
	//	response.FailWithMessage("invalid param "+key, c)
	//	c.Abort()
	//}
	return 0
}
