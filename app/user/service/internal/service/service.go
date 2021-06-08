package service

import (
	"github.com/google/wire"
	v1 "lucky-draw/api/user/service/v1"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewUserService)


type UserService struct {
	v1.UnimplementedUserServer
}

func NewUserService() *UserService {
	return &UserService{}
}