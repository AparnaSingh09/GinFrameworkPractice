package server

import "github.com/gin-gonic/gin"

func SetUpServer() *gin.Engine {
	server := gin.Default()

	server.GET("/test", func(context *gin.Context) {
		context.String(200, "endpoint working")
	})

	return server
}
