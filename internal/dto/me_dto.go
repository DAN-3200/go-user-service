package dto

type EditMeReq struct {
	ID       string `json:"id"`
	Name     string `json:"name" binding:"required,min=5"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=5,max=20"`
}
