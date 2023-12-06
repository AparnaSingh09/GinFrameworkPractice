package server

import (
	"GinFrameworkPractice/auth"
	"GinFrameworkPractice/controller"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var bookController *controller.BookController = controller.New()
var loginController *controller.LoginController = controller.NewLoginController()

func SetUpServer() *gin.Engine {
	server := gin.Default()
	//server.Use(basic_auth.BasicAuth())

	server.GET("/test", func(context *gin.Context) {
		context.String(200, "endpoint working")
	})

	// Login Endpoint: Authentication + Token creation
	server.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})

	apiRoutes := server.Group("/api", auth.AuthorizeJWT())
	{
		apiRoutes.GET("/books", func(context *gin.Context) {
			context.JSON(200, bookController.FindAll())
		})

		apiRoutes.GET("/books/:id", func(context *gin.Context) {
			idString := context.Param("id")
			id, _ := strconv.Atoi(idString)
			context.JSON(200, bookController.FindBookById(id))
		})

		apiRoutes.POST("/books", func(context *gin.Context) {
			context.JSON(200, bookController.AddBook(context))
		})

		apiRoutes.DELETE("/books", func(context *gin.Context) {
			idStr := context.Query("id")
			id, _ := strconv.Atoi(idStr)
			context.JSON(200, bookController.RemoveBook(id))
		})
	}

	return server
}
