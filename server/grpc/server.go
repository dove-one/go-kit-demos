package grpc

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"

	"net"
)

var opts = []grpc.ServerOption{
	grpc_middleware.WithUnaryServerChain(
		RecoveryInterceptor,
	),
}

var grpcServer = grpc.NewServer(opts...)

func RegisterService(ad *grpc.ServiceDesc, svc interface{}) {
	grpcServer.RegisterService(ad, svc)
}

func Run(addr string, errc chan error) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		errc <- err
		return
	}

	errc <- grpcServer.Serve(lis)
}
