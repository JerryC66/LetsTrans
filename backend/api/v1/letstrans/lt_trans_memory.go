package letstrans

import (
	"github.com/firwoodlin/letstrans/model/common/response"
	"github.com/gin-gonic/gin"
)

type TranslationMemoryApi struct{}

func (tma *TranslationMemoryApi) GetMemBySrc(c *gin.Context) {
	//jwtID := utils.GetUserID(c)
	srcText := c.Query("source_text")
	//srcLang := c.Query("source_lang")
	//targetLang := c.Query("target_lang")
	if srcText == "" {
		response.FailWithMessage("source_text is invalid", c)
		return
	}
	mem, err := transMemoryService.GetSuggestions(srcText)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(gin.H{"memories": mem}, c)
}
