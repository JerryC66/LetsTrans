package middleware

import (
	"bytes"
	"github.com/firwoodlin/letstrans/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io"
	"time"
)

func LogDetail() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 记录开始时间
		startTime := time.Now()

		// 记录请求方法和URL
		method := c.Request.Method
		url := c.Request.URL.String()

		// 记录请求头
		//headers := c.Request.Header
		clientIP := c.ClientIP()

		// 记录请求体
		bodyType := c.ContentType()
		var requestBody bytes.Buffer
		if c.Request.Body != nil {
			body, err := io.ReadAll(c.Request.Body)
			if err == nil {
				requestBody.Write(body)
				// 重新填充 Request.Body 以便后续处理
				c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
			}
		}

		// 处理请求
		c.Next()

		// 记录响应状态码和结束时间
		statusCode := c.Writer.Status()
		endTime := time.Now()
		duration := endTime.Sub(startTime)

		// 打印日志
		global.GVA_LOG.Info("Request details",
			zap.String("[client_ip]", clientIP),
			zap.String("[method]", method),
			zap.String("[url]", url),
			zap.Int("[status]", statusCode),
			zap.Duration("[duration]", duration),
			//zap.Any("[headers]", headers),
			zap.String("[bodyType]", bodyType),

			zap.String("[body]", requestBody.String()),
		)
	}
}
