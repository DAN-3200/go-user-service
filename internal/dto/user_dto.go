package dto

type UserReq struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=5,max=20"`
	Role     string `json:"role" binding:"required,oneof=admin user"`
}

type UserRes struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	CreatedAt string `json:"createAt"`
}

type UserUpdateReq struct {
	ID           string `json:"id" binding:"required"`
	Name         string `json:"name" binding:"required,min=5,max=20"`
	PasswordHash string `json:"password" binding:"required"`
}
