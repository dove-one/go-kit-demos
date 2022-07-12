package grpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"helloKit/endpoint"
	"helloKit/pb"
	svc "helloKit/service"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	model "helloKit/model"
)

type UserServer interface {
	Create(ctx context.Context, req *pb.CreateReq) (*pb.CreateResp, error)
	Delete(ctx context.Context, req *pb.DeleteReq) (*pb.DeleteResp, error)
}

type userServer struct {
	create grpctransport.Handler
	delete grpctransport.Handler
}

func NewUserServer(svc svc.UserSvc, opts ...grpctransport.ServerOption) UserServer {
	return userServer{
		create: grpctransport.NewServer(
			endpoint.MakeCreateEndpoint(svc),
			decodeCreateRequest,
			encodeCreateResponse,
			opts...,
		),
		delete: grpctransport.NewServer(
			endpoint.MakeDeleteEndpoint(svc),
			decodeDeleteRequest,
			encodeDeleteResponse,
			opts...,
		),
	}
}

func (u userServer) Create(ctx context.Context, req *pb.CreateReq) (*pb.CreateResp, error) {
	_, rep, err := u.create.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateResp), nil
}

func (u userServer) Delete(ctx context.Context, req *pb.DeleteReq) (*pb.DeleteResp, error) {
	_, rep, err := u.delete.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.DeleteResp), nil
}

func decodeCreateRequest(c context.Context, grpcReq interface{}) (interface{}, error) {
	req, ok := grpcReq.(*pb.CreateReq)
	if !ok {
		fmt.Println("haha")
		return nil, fmt.Errorf("grpc server decode request出错！")
	}
	request := &model.CreateReq{
		Name: req.Name,
		Age:  req.Age,
	}
	return request, nil
}

func encodeCreateResponse(c context.Context, response interface{}) (interface{}, error) {
	resp, ok := response.(*model.CreateResp)
	if !ok {
		return nil, fmt.Errorf("grpc server encode response error (%T)", response)
	}
	r := &pb.CreateResp{
		Code: resp.Code,
		Msg:  resp.Msg,
		Data: &pb.CreateRespData{
			Id:   resp.Data.Id,
			Age:  resp.Data.Age,
			Name: resp.Data.Name,
		},
	}
	return r, nil
}

//TODO 这里的逻辑要再看
func MakeCreateHandler(fullMethod string) func(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	return func(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
		in := new(pb.CreateReq)
		if err := dec(in); err != nil {
			return nil, err
		}
		if interceptor == nil {
			return srv.(UserServer).Create(ctx, in)
		}
		info := &grpc.UnaryServerInfo{
			Server:     srv,
			FullMethod: fullMethod,
		}
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.(UserServer).Create(ctx, req.(*pb.CreateReq))
		}
		return interceptor(ctx, in, info, handler)
	}
}

// Server
// 1. decode request          pb -> model
func decodeDeleteRequest(c context.Context, grpcReq interface{}) (interface{}, error) {
	req, ok := grpcReq.(*pb.DeleteReq)
	if !ok {
		return nil, fmt.Errorf("grpc server decode request出错！")
	}
	request := &model.DeleteReq{
		Name: req.Name,
		Id:   req.Id,
	}
	return request, nil
}

// 2. encode response           model -> pb
func encodeDeleteResponse(c context.Context, response interface{}) (interface{}, error) {
	resp, ok := response.(*model.DeleteResp)
	if !ok {
		return nil, fmt.Errorf("grpc server encode response出错！")
	}
	r := &pb.DeleteResp{
		Code: resp.Code,
		Msg:  resp.Msg,
		Data: &pb.DeleteRespData{
			Result: resp.Data.Result,
		},
	}
	return r, nil
}

func MakeDeleteHandler(fullMethod string) func(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	return func(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
		in := new(pb.DeleteReq)
		if err := dec(in); err != nil {
			return nil, err
		}
		if interceptor == nil {
			return srv.(UserServer).Delete(ctx, in)
		}
		info := &grpc.UnaryServerInfo{
			Server:     srv,
			FullMethod: fullMethod,
		}
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.(UserServer).Delete(ctx, req.(*pb.DeleteReq))
		}
		return interceptor(ctx, in, info, handler)
	}
}
