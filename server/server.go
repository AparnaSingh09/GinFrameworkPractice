package server

import (
	"GinFrameworkPractice/controller"
	"github.com/gin-gonic/gin"
)

var bookController *controller.BookController = controller.New()

func SetUpServer() *gin.Engine {
	server := gin.Default()

	server.GET("/test", func(context *gin.Context) {
		context.String(200, "endpoint working")
	})

	server.GET("/books", func(context *gin.Context) {
		context.JSON(200, bookController.FindAll())
	})

	return server
}
