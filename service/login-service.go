package service

type LoginService struct {
	authorizedUsername string
	authorizedPassword string
}

func NewLoginService() *LoginService {
	return &LoginService{
		authorizedUsername: "TW",
		authorizedPassword: "tw",
	}
}

func (loginService *LoginService) Login(username string, password string) bool {
	return false
}
