package server

import (
	basic_auth "GinFrameworkPractice/basic-auth"
	"GinFrameworkPractice/controller"
	"github.com/gin-gonic/gin"
	"strconv"
)

var bookController *controller.BookController = controller.New()

func SetUpServer() *gin.Engine {
	server := gin.Default()
	server.Use(basic_auth.BasicAuth())

	server.GET("/test", func(context *gin.Context) {
		context.String(200, "endpoint working")
	})

	server.GET("/books", func(context *gin.Context) {
		context.JSON(200, bookController.FindAll())
	})

	server.GET("/books/:id", func(context *gin.Context) {
		idString := context.Param("id")
		id, _ := strconv.Atoi(idString)
		context.JSON(200, bookController.FindBookById(id))
	})

	server.POST("/books", func(context *gin.Context) {
		context.JSON(200, bookController.AddBook(context))
	})

	server.DELETE("/books", func(context *gin.Context) {
		idStr := context.Query("id")
		id, _ := strconv.Atoi(idStr)
		context.JSON(200, bookController.RemoveBook(id))
	})

	return server
}
