
package service

import(
	"context"

	v1 "lucky-draw/api/user/service/v1"
)

func (s *UserService) CreateUser(ctx context.Context, req *v1.CreateUserReq) (*v1.CreateUserReply, error) {
	return &v1.CreateUserReply{}, nil
}
func (s *UserService) UpdateUser(ctx context.Context, req *v1.UpdateUserReq) (*v1.UpdateUserReply, error) {
	return &v1.UpdateUserReply{}, nil
}
func (s *UserService) GetUser(ctx context.Context, req *v1.GetUserReq) (*v1.GetUserReply, error) {
	return &v1.GetUserReply{}, nil
}
