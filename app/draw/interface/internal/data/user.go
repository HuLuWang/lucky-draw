package data

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"lucky-draw/app/draw/interface/internal/biz"
	
	usV1 "lucky-draw/api/user/service/v1"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/user")),
	}
}

func (rp *userRepo) Register(ctx context.Context, u *biz.User) (*biz.User, error) {
	rv, err := rp.data.uc.CreateUser(ctx, &usV1.CreateUserReq{
		Nickname: u.Nickname,
		Mobile:   u.Mobile,
		Password: u.Password,
	})
	
	return &biz.User{
		Id: rv.Id,
	}, err
}

// todo
func (rp *userRepo) Login(ctx context.Context, u *biz.User) (string, error) {
	rv, err := rp.data.uc.VerifyPassword(ctx, &usV1.VerifyPasswordReq{
		Mobile:   u.Mobile,
		Password: u.Password,
	})
	if err != nil {
		return "", err
	}
	if rv.Ok {
		return "token", nil
	}
	return "", errors.New("login failed")
}

func (rp *userRepo) Logout(ctx context.Context, u *biz.User) error {
	return nil
}
