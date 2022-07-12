package grpc

import (
	server "helloKit/server/grpc"
	svc "helloKit/service"
	transport "helloKit/transport/grpc"

	"google.golang.org/grpc"
)

func init() {
	registerUserService()
}

func registerUserService() {
	server.RegisterService(&grpc.ServiceDesc{
		ServiceName: "pb.User",
		HandlerType: (*transport.UserServer)(nil),
		Metadata:    "user.proto",
		Methods: []grpc.MethodDesc{
			{
				MethodName: "Create",
				Handler:    transport.MakeCreateHandler("/pb.User/Create"),
			},
			{
				MethodName: "Delete",
				Handler:    transport.MakeDeleteHandler("/pb.User/Delete"),
			},
		},
	}, transport.NewUserServer(svc.NewUserSvc()))
}
