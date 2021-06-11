package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"lucky-draw/app/user/service/internal/biz"
	"lucky-draw/app/user/service/internal/data/ent/user"
	"lucky-draw/app/user/service/internal/pkg/util"
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
	ph, err := util.HashPassword(u.Password)
	if err != nil {
		return nil, err
	}
	po, err := r.data.db.User.
		Create().
		SetNickname(u.Nickname).
		SetAvatar(u.Avatar).
		SetPasswordHash(ph).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.User{Id: po.ID, Nickname: po.Nickname, Avatar: po.Avatar}, nil
}

func (r *userRepo) GetUser(ctx context.Context, id int64) (*biz.User, error) {
	po, err := r.data.db.User.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return &biz.User{Id: po.ID, Nickname: po.Nickname, Avatar: po.Avatar}, nil
}

func (r *userRepo) UpdateUser(ctx context.Context, u *biz.User) (*biz.User, error) {
	po, err := r.data.db.User.UpdateOneID(u.Id).
		SetAvatar(u.Avatar).
		SetNickname(u.Nickname).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.User{
		Id:       po.ID,
		Nickname: po.Nickname,
		Avatar:   po.Avatar,
	}, nil
}

func (r *userRepo) VerifyPassword(ctx context.Context, u *biz.User) (bool, error) {
	po, err := r.data.db.User.
		Query().
		Where(user.MobileEQ(u.Mobile)).
		Only(ctx)
	if err != nil {
		return false, err
	}
	return util.CheckPasswordHash(u.Password, po.PasswordHash), nil
}
