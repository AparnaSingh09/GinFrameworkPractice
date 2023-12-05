package controller

import (
	"GinFrameworkPractice/entity"
	"GinFrameworkPractice/service"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	loginService *service.LoginService
	jwtService   *service.JWTService
}

func NewLoginController() *LoginController {
	return &LoginController{
		loginService: service.NewLoginService(),
		jwtService:   service.NewJWTService(),
	}
}

func (loginController *LoginController) Login(context *gin.Context) string {
	var credentials entity.Credentials
	context.BindJSON(&credentials)
	isAuthenticated := loginController.loginService.Login(credentials.Username, credentials.Password)
	if isAuthenticated {
		return loginController.jwtService.GenerateToken(credentials.Username)
	}
	return ""
}
