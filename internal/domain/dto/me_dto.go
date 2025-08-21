package dto

type EditMeReq struct {
	Name     *string `json:"name" binding:"omitempty,min=5,max=20"`
	IsActive *bool   `json:"isActive" binding:"omitempty,boolean"`
}
