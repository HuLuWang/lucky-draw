package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"lucky-draw/app/user/service/internal/biz"
	"lucky-draw/app/user/service/internal/data/ent/address"
)

var _ biz.AddressRepo = (*addressRepo)(nil)

type addressRepo struct {
	data *Data
	log  *log.Helper
}

func (r *addressRepo) CreateAddress(ctx context.Context, a *biz.Address) (*biz.Address, error) {
	po, err := r.data.db.Address.
		Create().
		SetUserID(a.UserId).
		SetAddress(a.Address).
		SetMobile(a.Mobile).
		SetUsername(a.Username).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Address{
		Id:       po.ID,
		UserId:   po.UserID,
		Username: po.Username,
		Mobile:   po.Mobile,
		Address:  po.Address,
	}, nil
}

func (r *addressRepo) GetAddress(ctx context.Context, id int64) (*biz.Address, error) {
	po, err := r.data.db.Address.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return &biz.Address{
		Id:       po.ID,
		UserId:   po.UserID,
		Username: po.Username,
		Mobile:   po.Mobile,
		Address:  po.Address,
	}, nil
}

func (r *addressRepo) UpdateAddress(ctx context.Context, address *biz.Address) (*biz.Address, error) {
	po, err := r.data.db.Address.UpdateOneID(address.Id).
		SetUsername(address.Username).
		SetMobile(address.Mobile).
		SetAddress(address.Address).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Address{
		Id:       po.ID,
		UserId:   po.UserID,
		Username: po.Username,
		Mobile:   po.Mobile,
		Address:  po.Address,
	}, nil
}

func (r *addressRepo) DeleteAddress(ctx context.Context, id int64) (bool, error) {
	err := r.data.db.Address.DeleteOneID(id).Exec(ctx)
	return err == nil, err
}

func (r *addressRepo) ListAddress(ctx context.Context, uid int64) ([]*biz.Address, error) {
	pos, err := r.data.db.Address.Query().Where(address.UserIDEQ(uid)).All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Address, 0)
	for _, po := range pos {
		rv = append(rv, &biz.Address{
			Id:       po.ID,
			UserId:   po.UserID,
			Username: po.Username,
			Mobile:   po.Mobile,
			Address:  po.Address,
		})
	}
	return rv, nil
}