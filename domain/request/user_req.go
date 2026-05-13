package request

type LoginRequest struct {
	Email    string `binding:"required" json:"email"`
	Password string `binding:"required" json:"password"`
}

type RegistrationRequest struct {
	Name     string `binding:"required" json:"name"`
	Email    string `binding:"required" json:"email"`
	Phone    string `json:"phone"`
	Password string `binding:"required" json:"password"`
	Status   string `json:"status"`
}
