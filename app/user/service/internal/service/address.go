package service

import (
	"context"
	v1 "lucky-draw/api/user/service/v1"
)

func (s *UserService) CreateAddress(ctx context.Context, req *v1.CreateAddressReq) (*v1.CreateAddressReply, error) {
	return &v1.CreateAddressReply{}, nil
}
func (s *UserService) UpdateAddress(ctx context.Context, req *v1.UpdateAddressReq) (*v1.UpdateAddressReply, error) {
	return &v1.UpdateAddressReply{}, nil
}
func (s *UserService) DeleteAddress(ctx context.Context, req *v1.DeleteAddressReq) (*v1.DeleteAddressReply, error) {
	return &v1.DeleteAddressReply{}, nil
}
func (s *UserService) ListAddress(ctx context.Context, req *v1.ListAddressReq) (*v1.ListAddressReply, error) {
	return &v1.ListAddressReply{}, nil
}

