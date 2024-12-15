package mapper

import (
	"github.com/givxl33t/go-fiber-boilerplate/internal/domain"
	"github.com/givxl33t/go-fiber-boilerplate/internal/model"
)

func ToUserResponse(user *domain.User) *model.UserResponse {
	return &model.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
