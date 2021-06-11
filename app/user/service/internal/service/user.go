package service

import (
	"context"
	"lucky-draw/app/user/service/internal/biz"
	
	v1 "lucky-draw/api/user/service/v1"
)

func (s *UserService) CreateUser(ctx context.Context, req *v1.CreateUserReq) (*v1.CreateUserReply, error) {
	rv, err := s.uc.Create(ctx, &biz.User{
		Nickname: req.Nickname,
		Mobile:   req.Mobile,
		Password: req.Password,
	})
	return &v1.CreateUserReply{
		Id: rv.Id,
	}, err
}

func (s *UserService) UpdateUser(ctx context.Context, req *v1.UpdateUserReq) (*v1.UpdateUserReply, error) {
	rv, err := s.uc.Update(ctx, &biz.User{
		Id:       req.Id,
		Nickname: req.Nickname,
		Mobile:   req.Mobile,
		Avatar:   req.Avatar,
	})
	return &v1.UpdateUserReply{Res: &v1.UserInfo{
		Id:       rv.Id,
		Nickname: rv.Nickname,
		Mobile:   rv.Mobile,
		Avatar:   rv.Avatar,
	}}, err
}
func (s *UserService) GetUser(ctx context.Context, req *v1.GetUserReq) (*v1.GetUserReply, error) {
	rv, err := s.uc.Get(ctx, req.Id)
	return &v1.GetUserReply{Res: &v1.UserInfo{
		Id:       rv.Id,
		Nickname: rv.Nickname,
		Mobile:   rv.Mobile,
		Avatar:   rv.Avatar,
	}}, err
}

func (s *UserService) VerifyPassword(ctx context.Context, req *v1.VerifyPasswordReq) (*v1.VerifyPasswordReply, error) {
	rv, err := s.uc.VerifyPassword(ctx, &biz.User{Mobile: req.Mobile, Password: req.Password})
	return &v1.VerifyPasswordReply{
		Ok: rv,
	}, err
}
