package users

type SignUpRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" validate:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

type SignUpResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Success bool   `json:"success"`
	Token   string `json:"token,omitempty"`
	Error   error  `json:"error,omitempty"`
}
