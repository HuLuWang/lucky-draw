package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Id       int64
	Nickname string
	Mobile   string
	Password string
	Avatar   string
}

type UserRepo interface {
	Register(ctx context.Context, u *User) (*User, error)
	Login(ctx context.Context, u *User) (string, error)
	Logout(ctx context.Context, u *User) error
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	log := log.NewHelper(log.With(logger, "module", "usecase/interface"))
	return &UserUseCase{
		repo: repo,
		log:  log,
	}
}

func (uc *UserUseCase) Register(ctx context.Context, u *User) (*User, error) {
	return uc.repo.Register(ctx, u)
}

func (uc *UserUseCase) Login(ctx context.Context, u *User) (string, error) {
	return uc.repo.Login(ctx, u)
}

func (uc *UserUseCase) Logout(ctx context.Context, u *User) error {
	return uc.repo.Logout(ctx, u)
}
