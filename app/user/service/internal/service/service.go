package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	v1 "lucky-draw/api/user/service/v1"
	"lucky-draw/app/user/service/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewUserService)


type UserService struct {
	v1.UnimplementedUserServer
	
	uc *biz.UserUseCase
	ac *biz.AddressUseCase
	log *log.Helper
}

func NewUserService(uc *biz.UserUseCase, ac *biz.AddressUseCase, logger log.Logger) *UserService {
	return &UserService{
		uc: uc,
		ac: ac,
		log: log.NewHelper(log.With(logger, "module", "service/server-service")),
	}
}