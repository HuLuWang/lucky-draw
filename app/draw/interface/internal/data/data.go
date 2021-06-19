package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	"go.opentelemetry.io/otel/propagation"
	"lucky-draw/app/draw/interface/internal/conf"
	
	userv1 "lucky-draw/api/user/service/v1"
	consulAPI "github.com/hashicorp/consul/api"
	
	"github.com/go-kratos/kratos/v2/transport/grpc"
	consul "github.com/go-kratos/consul/registry"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"

)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserRepo)

// Data .
type Data struct {
	log *log.Helper
	uc  userv1.UserClient
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, uc userv1.UserClient) (*Data, error) {
	l := log.NewHelper(log.With(logger, "module", "data"))
	return &Data{log: l, uc: uc}, nil
}

func NewDiscovery(conf *conf.Registry) registry.Discovery {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli)
	return r
}

func NewUserServiceClient(r registry.Discovery, tp *tracesdk.TracerProvider) userv1.UserClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///draw.user.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			tracing.Client(
				tracing.WithTracerProvider(tp),
				tracing.WithPropagators(
					propagation.NewCompositeTextMapPropagator(propagation.Baggage{}, propagation.TraceContext{}),
					),
				),
				recovery.Recovery(),
			),
	)
	if err != nil {
		panic(err)
	}
	c := userv1.NewUserClient(conn)
	return c
}
