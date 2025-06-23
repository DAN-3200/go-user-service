// DTO (Data Transfer Object)
package dto

type Login struct {
	Email    string `json:"email" binding:"required,email,min=10,max=40"`
	Password string `json:"password" binding:"required,min=5,max=20"`
}

type UserRegisterReq struct {
	Name     string `json:"name" binding:"required,min=5,max=20"`
	Email    string `json:"email" binding:"required,email,min=10,max=40"`
	Password string `json:"password" binding:"required,min=5,max=20"`
}

type RefreshPassword struct {
	NewPassword string `json:"password" binding:"required,min=5,max=20"`
}
