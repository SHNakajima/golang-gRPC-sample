package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "test/sample/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// gRPCサーバーへの接続
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewUserServiceClient(conn)

	// コンテキストの設定
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// GetUser APIの呼び出し
	fmt.Println("-- GetUser API --")
	user, err := client.GetUser(ctx, &pb.GetUserRequest{UserId: 1})
	if err != nil {
		log.Fatalf("could not get user: %v", err)
	}
	fmt.Printf("User: %s (%s), Age: %d\n", user.Name, user.Email, user.Age)

	// ListUsers APIの呼び出し
	fmt.Println("\n-- ListUsers API --")
	listResp, err := client.ListUsers(ctx, &pb.ListUsersRequest{Page: 1, PageSize: 10})
	if err != nil {
		log.Fatalf("could not list users: %v", err)
	}
	fmt.Printf("Total users: %d\n", listResp.TotalCount)
	for i, u := range listResp.Users {
		fmt.Printf("%d. %s (%s), Age: %d\n", i+1, u.Name, u.Email, u.Age)
	}

	// CreateUser APIの呼び出し
	fmt.Println("\n-- CreateUser API --")
	newUser, err := client.CreateUser(ctx, &pb.CreateUserRequest{
		Name:  "鈴木一郎",
		Email: "ichiro@example.com",
		Age:   35,
	})
	if err != nil {
		log.Fatalf("could not create user: %v", err)
	}
	fmt.Printf("Created new user: ID=%d, Name=%s, Email=%s, Age=%d, CreatedAt=%s\n",
		newUser.Id, newUser.Name, newUser.Email, newUser.Age, newUser.CreatedAt)

	// 作成したユーザーの確認
	fmt.Println("\n-- Verify Created User --")
	createdUser, err := client.GetUser(ctx, &pb.GetUserRequest{UserId: newUser.Id})
	if err != nil {
		log.Fatalf("could not get created user: %v", err)
	}
	fmt.Printf("Retrieved user: ID=%d, Name=%s, Email=%s, Age=%d, CreatedAt=%s\n",
		createdUser.Id, createdUser.Name, createdUser.Email, createdUser.Age, createdUser.CreatedAt)
}