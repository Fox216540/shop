package user_dto

type UserLoginRequest struct {
	UsernameOrEmail string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}
type UserRegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=20"`
}
type UserRegisterResponse struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}
