package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"lucky-draw/app/user/service/internal/biz"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/user-service")),
	}
}


func (r *userRepo) CreateUser(ctx context.Context, u *biz.User) (*biz.User, error) {
}

func (r *userRepo) GetUser(ctx context.Context, id int64) (*biz.User, error) {
}

func (r *userRepo) UpdateUser(ctx context.Context, u *biz.User) (*biz.User, error) {
}

func (r *userRepo) VerifyPassword(ctx context.Context, u *biz.User) (bool, error) {
}


