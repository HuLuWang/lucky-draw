package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	v1 "lucky-draw/api/draw/interface/v1"
	"lucky-draw/app/draw/interface/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewDrawInterface)

type DrawInterface struct {
	v1.UnimplementedDrawInterfaceServer
	uc  *biz.UserUseCase
	log *log.Helper
}

func NewDrawInterface(uc *biz.UserUseCase, logger log.Logger) *DrawInterface {
	return &DrawInterface{
		log: log.NewHelper(log.With(logger, "module", "service/interface")),
		uc:  uc,
	}
}
