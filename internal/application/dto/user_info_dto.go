package dto

import (
	"app/internal/domain/entity"
	"time"
)

type EditMeReq struct {
	Name     *string `json:"name" binding:"omitempty,min=5,max=20"`
	IsActive *bool   `json:"isActive" binding:"omitempty,boolean"`
}

type UserInfoRes struct {
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	IsEmailVerified bool      `json:"isEmailVerified"`
	CreatedAt       time.Time `json:"createdAt"`
	Role            string    `json:"role"`
}

func ToUserResponse(u *entity.User) *UserInfoRes {
	return &UserInfoRes{
		ID:              u.ID(),
		Name:            u.Name(),
		Email:           u.Email(),
		IsEmailVerified: u.IsEmailVerified(),
		CreatedAt:       u.CreatedAt(),
		Role:            string(u.Role()),
	}
}
