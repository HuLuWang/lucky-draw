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
	CreateUser(ctx context.Context, u *User) (*User, error)
	GetUser(ctx context.Context, id int64) (*User, error)
	UpdateUser(ctx context.Context, u *User) (*User, error)
	VerifyPassword(ctx context.Context, u *User) (bool, error)
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "usercase/user")),
	}
}

func (uc *UserUseCase) Create(ctx context.Context, u *User) (*User, error) {
	return uc.repo.CreateUser(ctx, u)
}

func (uc *UserUseCase) Get(ctx context.Context, id int64) (*User, error) {
	return uc.repo.GetUser(ctx, id)
}

func (uc *UserUseCase) Update(ctx context.Context, u *User) (*User, error) {
	return uc.repo.UpdateUser(ctx, u)
}

func (uc *UserUseCase) VerifyPassword(ctx context.Context, u *User) (bool, error) {
	return uc.repo.VerifyPassword(ctx, u)
}
