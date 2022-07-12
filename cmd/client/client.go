package main

import (
	"context"
	"flag"
	"fmt"
	"helloKit/conf"
	"time"

	"google.golang.org/grpc"

	pb "helloKit/pb"
)

func init() {
	// flags
	flag.StringVar(&conf.GrpcAddr, "grpc-addr", conf.GetEnv("GrpcAddr", "0.0.0.0:5000"), "grpc服务地址")
}

func main() {
	conn, err := grpc.Dial(conf.GrpcAddr, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("创建grpc连接失败! Error: %s", err)
	}
	defer conn.Close()

	Create(conn)
	//	Delete(conn)
}

func Create(conn *grpc.ClientConn) {
	in := &pb.CreateReq{
		Name: "wss",
		Age:  19,
	}

	out := &pb.CreateResp{}
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	err := conn.Invoke(ctx, "/pb.User/Create", in, out)

	fmt.Println(err)
	fmt.Println(out)
}
func Delete(conn *grpc.ClientConn) {
	in := &pb.DeleteReq{
		Name: "wss",
		Id:   19,
	}

	out := &pb.DeleteResp{}
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	err := conn.Invoke(ctx, "/pb.User/Delete", in, out)
	fmt.Println(err)
	fmt.Println(out)
}
