package dto

type UserReq struct {
	Name     string `json:"name" binding:"required,min=5,max=20"`
	Email    string `json:"email" binding:"required,email,min=10,max=40"`
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

type EditUserReq struct {
	Name            *string `json:"name" binding:"omitempty,min=5,max=20"`
	Email           *string `json:"email" binding:"omitempty,email,min=10,max=40"`
	Password        *string `json:"password" binding:"omitempty,min=5,max=20"`
	IsEmailVerified *bool   `json:"isEmailVerified" binding:"omitempty,boolean"`
	IsActive        *bool   `json:"isActive" binding:"omitempty,boolean"`
	Role            *string `json:"role" binding:"omitempty,oneof=admin user"`
}
