package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"

	proto "github.com/pudongping/golang-grpc-demo/proto/user"
)

const address = "0.0.0.0:8081"

func userIndex(ctx context.Context, userClient proto.UserClient) {
	// UserIndex 请求
	userIndexResponse, err := userClient.UserIndex(ctx, &proto.UserIndexRequest{
		Page:     2,
		PageSize: 5,
	})
	if err != nil {
		log.Printf("user index err ==> %v", err)
	}

	if 0 == userIndexResponse.Code {
		log.Printf("user index success: %s", userIndexResponse.Msg)
		userEntityList := userIndexResponse.GetData()
		for _, row := range userEntityList {
			fmt.Println("list ==> name ==> ", row.Name, " age ==> ", row.Age)
		}
	} else {
		log.Printf("user index error: %d", userIndexResponse.Code)
	}
}

func userView(ctx context.Context, userClient proto.UserClient) {
	userViewResponse, err := userClient.UserView(ctx, &proto.UserViewRequest{Uid: 1})
	if err != nil {
		log.Printf("user view err ==> %v", err)
	}

	if 0 == userViewResponse.Code {
		log.Printf("user view success: %s", userViewResponse.Msg)
		userEntity := userViewResponse.Data
		fmt.Println("view ==> name ==> ", userEntity.Name, " age ==> ", userEntity.Age)
	} else {
		log.Printf("user view error: %d", userViewResponse.Code)
	}
}

func userPost(ctx context.Context, userClient proto.UserClient) {
	userPostResponse, err := userClient.UserPost(ctx, &proto.UserPostRequest{
		Name:     "Alex",
		Password: "123456",
		Age:      27,
	})
	if err != nil {
		log.Printf("user post err ==> %v", err)
	}

	if 0 == userPostResponse.Code {
		log.Printf("user post success: %s", userPostResponse.Msg)
	} else {
		log.Printf("user post error: %d", userPostResponse.Code)
	}
}

func userDelete(ctx context.Context, userClient proto.UserClient) {
	userDeleteResponse, err := userClient.UserDelete(ctx, &proto.UserDeleteRequest{Uid: 2})
	if err != nil {
		log.Printf("user delete err ==> %v", err)
	}

	if 0 == userDeleteResponse.Code {
		log.Printf("user delete success: %s", userDeleteResponse.Msg)
	} else {
		log.Printf("user delete error: %d", userDeleteResponse.Code)
	}
}

func main() {
	// 建立连接

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("does not connect: %v", err)
	}

	defer conn.Close()

	userClient := proto.NewUserClient(conn)

	// 设定请求超时时间为 3s
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	userIndex(ctx, userClient)
	userView(ctx, userClient)
	userPost(ctx, userClient)
	userDelete(ctx, userClient)

}
