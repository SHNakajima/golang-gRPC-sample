package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	pb "test/sample/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedUserServiceServer
	users map[int32]*pb.User
}

func newServer() *server {
	s := &server{
		users: make(map[int32]*pb.User),
	}
	// サンプルデータの追加
	s.users[1] = &pb.User{
		Id:        1,
		Name:      "山田太郎",
		Email:     "taro@example.com",
		Age:       30,
		CreatedAt: time.Now().Format(time.RFC3339),
	}
	s.users[2] = &pb.User{
		Id:        2,
		Name:      "佐藤花子",
		Email:     "hanako@example.com",
		Age:       25,
		CreatedAt: time.Now().Format(time.RFC3339),
	}
	return s
}

func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	// ユーザーIDのバリデーション
	if req.UserId <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid user ID: %v", req.UserId)
	}

	// ユーザーの取得
	user, exists := s.users[req.UserId]
	if !exists {
		return nil, status.Errorf(codes.NotFound, "User not found: %v", req.UserId)
	}

	return user, nil
}

func (s *server) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	// ページングパラメータのバリデーション
	if req.Page < 1 {
		req.Page = 1
	}
	if req.PageSize < 1 || req.PageSize > 100 {
		req.PageSize = 10
	}

	// すべてのユーザーを配列に変換
	allUsers := make([]*pb.User, 0, len(s.users))
	for _, user := range s.users {
		allUsers = append(allUsers, user)
	}

	// 簡易的なページング処理（実際のアプリケーションではより複雑な処理が必要）
	start := (req.Page - 1) * req.PageSize
	end := start + req.PageSize
	if start >= int32(len(allUsers)) {
		return &pb.ListUsersResponse{
			Users:      []*pb.User{},
			TotalCount: int32(len(allUsers)),
		}, nil
	}
	if end > int32(len(allUsers)) {
		end = int32(len(allUsers))
	}

	return &pb.ListUsersResponse{
		Users:      allUsers[start:end],
		TotalCount: int32(len(allUsers)),
	}, nil
}

func (s *server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.User, error) {
	// リクエストのバリデーション
	if req.Name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Name is required")
	}
	if req.Email == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Email is required")
	}

	// 新しいユーザーIDの生成
	newID := int32(len(s.users) + 1)
	for s.users[newID] != nil {
		newID++
	}

	// 新しいユーザーの作成
	user := &pb.User{
		Id:        newID,
		Name:      req.Name,
		Email:     req.Email,
		Age:       req.Age,
		CreatedAt: time.Now().Format(time.RFC3339),
	}

	// ユーザーの保存
	s.users[newID] = user

	return user, nil
}

func main() {
	// gRPCサーバーの起動準備
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, newServer())

	fmt.Println("Starting gRPC server on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}