package rpcserver

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"chative-server-go/rpcserver/internal/config"
	"chative-server-go/rpcserver/internal/server"
	"chative-server-go/rpcserver/internal/svc"
	"chative-server-go/rpcserver/pb"

	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Rpcserver(conf zrpc.RpcServerConf) {

	var c = config.Config{RpcServerConf: conf}
	ctx := svc.NewServiceContext(c)

	rand.Seed(time.Now().Unix())

	// zrpc.DontLogContentAlways(true)
	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterFriendServer(grpcServer, server.NewFriendServer(ctx))
		pb.RegisterRegiterServer(grpcServer, server.NewRegiterServer(ctx))
		pb.RegisterAccountServer(grpcServer, server.NewAccountServer(ctx))
		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	s.AddUnaryInterceptors(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		return handler(ctx, req)
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
