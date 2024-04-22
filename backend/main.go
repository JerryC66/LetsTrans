package main

import "github.com/gin-gonic/gin"
import "github.com/firwoodlin/letstrans/controller"

func main() {
	r := gin.Default()
	r.GET("/ping", controller.Ping)
	_ = r.Run()
}
