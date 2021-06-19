package service

import (
	"context"
	v1 "lucky-draw/api/draw/interface/v1"
	"lucky-draw/app/draw/interface/internal/biz"
)

func (d *DrawInterface) Register(ctx context.Context, req *v1.RegisterReq) (*v1.RegisterReply, error) {
	rv, err := d.uc.Register(ctx, &biz.User{
		Mobile:   req.Mobile,
		Password: req.Password,
	})
	return &v1.RegisterReply{
		Id: rv.Id,
	}, err
}

func (d *DrawInterface) Login(ctx context.Context, req *v1.LoginReq) (*v1.LoginReply, error) {
	rv, err := d.uc.Login(ctx, &biz.User{
		Mobile:   req.Mobile,
		Password: req.Password,
	})
	return &v1.LoginReply{
		Token: rv,
	}, err
}

func (d *DrawInterface) Logout(ctx context.Context, req *v1.LogoutReq) (*v1.LogoutReply, error) {
	err := d.uc.Logout(ctx, &biz.User{})
	return &v1.LogoutReply{}, err
}
