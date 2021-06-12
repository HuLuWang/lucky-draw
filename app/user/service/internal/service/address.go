package service

import (
	"context"
	v1 "lucky-draw/api/user/service/v1"
	"lucky-draw/app/user/service/internal/biz"
)

func (s *UserService) CreateAddress(ctx context.Context, req *v1.CreateAddressReq) (*v1.CreateAddressReply, error) {
	rv, err := s.ac.Create(ctx, &biz.Address{
		UserId:   req.Uid,
		Username: req.Username,
		Mobile:   req.Mobile,
		Address:  req.Address,
	})
	return &v1.CreateAddressReply{
		Id: rv.Id,
	}, err
}

func (s *UserService) UpdateAddress(ctx context.Context, req *v1.UpdateAddressReq) (*v1.UpdateAddressReply, error) {
	rv, err := s.ac.Update(ctx, &biz.Address{
		Id:       req.Id,
		Username: req.Username,
		Mobile:   req.Mobile,
		Address:  req.Address,
	})
	return &v1.UpdateAddressReply{
		Address: &v1.Address{
			Id:       rv.Id,
			Uid:      rv.UserId,
			Username: rv.Username,
			Mobile:   rv.Mobile,
			Address:  rv.Address,
		},
	}, err
}
func (s *UserService) DeleteAddress(ctx context.Context, req *v1.DeleteAddressReq) (*v1.DeleteAddressReply, error) {
	rv, err := s.ac.Delete(ctx, req.Id)
	return &v1.DeleteAddressReply{Ok: rv}, err
}
func (s *UserService) ListAddress(ctx context.Context, req *v1.ListAddressReq) (*v1.ListAddressReply, error) {
	rv, err := s.ac.List(ctx, req.Uid)
	rs := make([]*v1.Address, 0)
	for _, x := range rv {
		rs = append(rs, &v1.Address{
			Id:       x.Id,
			Uid:      x.UserId,
			Username: x.Username,
			Mobile:   x.Mobile,
			Address:  x.Address,
		})
	}
	return &v1.ListAddressReply{Results: rs}, err
}
