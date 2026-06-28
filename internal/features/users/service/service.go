package user_service

import (
	"context"

	"github.com/Sinhofazatron/tasks-go/internal/core/domain"
)

type UsersService struct {
	usersRepository UserRepository
}

type UserRepository interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	GetUser(ctx context.Context, id int) (domain.User, error)
	PatchUser(ctx context.Context, id int, user domain.User) (domain.User, error)
	DeleteUser(ctx context.Context, id int) error
	GetUsers(ctx context.Context, limit, offset *int) ([]domain.User, error)
}

func NewUserService(userRepository UserRepository) *UsersService {
	return &UsersService{
		usersRepository: userRepository,
	}
}
