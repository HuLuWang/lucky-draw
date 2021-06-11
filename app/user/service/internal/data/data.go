package data

import (
	"context"
	
	"lucky-draw/app/user/service/internal/data/ent"
	"lucky-draw/app/user/service/internal/conf"
	
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	
	_ "github.com/go-sql-driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserRepo, NewAddressRepo)

// Data .
type Data struct {
	db *ent.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "server-service/user-data"))
	
	client, err := ent.Open(c.Database.Driver, c.Database.Driver)
	if err != nil {
		log.Errorf("fail to open sql: %v", err)
		return nil, nil, err
	}
	// run auto migration tool
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Errorf("fail create schema resource %v", err)
		return nil, nil, err
	}
	
	d := &Data{
		db: client,
	}
	
	return d, func() {
		if err := d.db.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}
