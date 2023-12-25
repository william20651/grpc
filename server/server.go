package main

import (
 "context"
 "log"
 "net"
 "os"
 users "user-service/service" // 导入之前生成的包

 "google.golang.org/grpc"
)

// userService类型是Users服务的服务处理程序
type userService struct {
 users.UnimplementedUsersServer // 这个字段对于gRPC中的任何服务实现都是强制性的
}

func (s *userService) GetUser(ctx context.Context, in *users.UserGetRequest) (*users.UserGetReply, error) {
 // 打印客户端传过来的数据
 log.Printf("已接收到邮件地址: %s, 还有ID: %d", in.Email, in.Id)

 // 自定义数据响应给客户端
 u := users.User{
  Id:        "user-782935",
  FirstName: "tan",
  LastName:  "tianran",
  Age:       30,
 }
 return &users.UserGetReply{User: &u}, nil
}

// 向gRPC服务器注册Users服务
func registerServices(s *grpc.Server) {
 users.RegisterUsersServer(s, &userService{})
}

// 启动gRPC服务器
func startServer(s *grpc.Server, l net.Listener) error {
 return s.Serve(l)
}

func main() {
 listenAddr := os.Getenv("LISTEN_ADDR")

 if len(listenAddr) == 0 {
  listenAddr = ":50051"
 }

 lis, err := net.Listen("tcp", listenAddr)

 if err != nil {
  log.Fatal(err)
 }

 s := grpc.NewServer()
 registerServices(s)

 log.Fatal(startServer(s, lis))
}