package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Address struct {
	Id int64
	UserId int64
	Username string
	Mobile string
	Address string
}

type AddressRepo interface {
	CreateAddress(ctx context.Context, u *Address) (*Address, error)
	GetAddress(ctx context.Context, id int64) (*Address, error)
	UpdateAddress(ctx context.Context, u *Address) (*Address, error)
	DeleteAddress(ctx context.Context, id int64) (bool, error)
	ListAddress(ctx context.Context, uid int64) ([]*Address, error)
}

type AddressUseCase struct {
	repo AddressRepo
	log  *log.Helper
}

func NewAddressUseCase(repo AddressRepo, logger log.Logger) *AddressUseCase {
	return &AddressUseCase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "usercase/address")),
	}
}

func (ac *AddressUseCase) Create(ctx context.Context, a *Address) (*Address, error) {
	return ac.repo.CreateAddress(ctx, a)
}

func (ac *AddressUseCase) Get(ctx context.Context, id int64) (*Address, error) {
	return ac.repo.GetAddress(ctx, id)
}

func (ac *AddressUseCase) Update(ctx context.Context, a *Address) (*Address, error) {
	return ac.repo.UpdateAddress(ctx, a)
}

func (ac *AddressUseCase) Delete(ctx context.Context, id int64) (bool, error) {
	return ac.repo.DeleteAddress(ctx, id)
}

func (ac *AddressUseCase) List(ctx context.Context, uid int64) ([]*Address, error) {
	return ac.repo.ListAddress(ctx, uid)
}


