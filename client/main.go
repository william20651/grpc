package main

import (
 "context"
 "log"
 "os"
 users "user-service/service" // 导入之前生成的包

 "google.golang.org/grpc"
)

// 建立与服务器的连接（通道）
func setupGrpcConnection(addr string) (*grpc.ClientConn, error) {
 return grpc.DialContext(
  context.Background(),
  addr,
  grpc.WithInsecure(),
  grpc.WithBlock(),
 )
}

// 创建客户端与Users服务通信
func getUserServiceClient(conn *grpc.ClientConn) users.UsersClient {
 return users.NewUsersClient(conn)
}

// 调用Users服务中的GetUser()方法
func getUser(client users.UsersClient, u *users.UserGetRequest) (*users.UserGetReply, error) {
 return client.GetUser(context.Background(), u)
}

func main() {
 if len(os.Args) != 2 {
  log.Fatal("缺少gRPC服务器地址")
 }
 conn, err := setupGrpcConnection(os.Args[1])
 if err != nil {
  log.Fatal(err)
 }
 defer conn.Close()
 c := getUserServiceClient(conn)
 result, err := getUser(c, &users.UserGetRequest{
  Email: "tantianran@qq.com",
  Id:    801896,
 })
 if err != nil {
  log.Fatal(err)
 }

 // 打印响应
 log.Printf("收到响应: %s %s %s %d\n", result.User.Id, result.User.FirstName, result.User.LastName, result.User.Age)
}