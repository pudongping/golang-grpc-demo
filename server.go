package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	proto "github.com/pudongping/golang-grpc-demo/proto/user"
)

const port = ":8081"

type UserService struct {
}

func (u *UserService) UserIndex(ctx context.Context, in *proto.UserIndexRequest) (*proto.UserIndexResponse, error) {
	log.Printf("receive user index request: page %d page_size %d", in.GetPage(), in.GetPageSize())

	user := []*proto.UserEntity{}
	user = append(user, &proto.UserEntity{
		Name: "Alex",
		Age:  26,
	})
	user = append(user, &proto.UserEntity{
		Name: "Harry",
		Age:  18,
	})

	return &proto.UserIndexResponse{
		Code: 0,
		Msg:  "Success",
		Data: user,
	}, nil
}

func (u *UserService) UserView(ctx context.Context, in *proto.UserViewRequest) (*proto.UserViewResponse, error) {
	log.Printf("receive user view request: uid %d", in.Uid)

	return &proto.UserViewResponse{
		Code: 0,
		Msg:  "Success",
		Data: &proto.UserEntity{Name: "Alex", Age: 26},
	}, nil
}

func (u *UserService) UserPost(ctx context.Context, in *proto.UserPostRequest) (*proto.UserPostResponse, error) {
	log.Printf("receive user post request: name %s, password %s, age %d", in.Name, in.Password, in.Age)

	return &proto.UserPostResponse{
		Code: 0,
		Msg:  "Success",
	}, nil
}

func (u *UserService) UserDelete(ctx context.Context, in *proto.UserDeleteRequest) (*proto.UserDeleteResponse, error) {
	log.Printf("receive user delete request: uid %d", in.Uid)

	return &proto.UserDeleteResponse{
		Code: 0,
		Msg:  "Success",
	}, nil
}

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		log.Printf("grpc server is start 0.0.0.0%s", port)
	}

	// 创建 GRPC 服务容器
	grpcServer := grpc.NewServer()

	// 为 User 服务注册业务实现（将 User 服务绑定到 GRPC 服务容器上）
	proto.RegisterUserServer(grpcServer, &UserService{})

	// 注册反射服务，这个服务是 CLI 使用的，跟服务本身没有关系
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
