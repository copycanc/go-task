package auth

type HandlerAuth struct {
	authService *AuthService
}

func NewHandlerAuth(authService *AuthService) *HandlerAuth {
	return &HandlerAuth{authService: authService}
}
