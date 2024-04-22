package controller

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

// TestPing 测试Ping函数
func TestPing(t *testing.T) {
	// 设置Gin的模式为测试模式
	gin.SetMode(gin.TestMode)

	// 创建一个Gin的记录器
	router := gin.New()
	router.GET("/ping", Ping)

	for count := 0; count <= 100000; count++ {
		// 使用httptest创建一个HTTP请求
		req, _ := http.NewRequest(http.MethodGet, "/ping", nil)
		// 使用httptest记录响应
		w := httptest.NewRecorder()
		// 将请求通过创建的路由处理
		router.ServeHTTP(w, req)

		// 检查HTTP状态码是否为200
		assert.Equal(t, http.StatusOK, w.Code)

		// 检查响应体是否为预期的JSON
		expected := `{"message":"pong"}`
		assert.JSONEq(t, expected, w.Body.String())
	}

}
