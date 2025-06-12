// DTO (Data Transfer Object)
package dto

type Login struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=5,max=20"`
}

type UserRegisterReq struct {
	Name     string `json:"name" binding:"required,min=5"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=5,max=20"`
}

type RefreshPassword struct {
	ID          string `json:"id,omitempty"`
	NewPassword string `json:"password" binding:"required,min=5,max=20"`
}
