package tests

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"testing"
	proto2 "xkshop/v1/xkshop_srv/user_srv/proto"
)

var userClient proto2.UserClient
var conn *grpc.ClientConn
var err error

//测试GetUserList接口

func TestGetUserList(t *testing.T) {
	conn, err = grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		panic(err)
	}
	userClient = proto2.NewUserClient(conn)

	rsp, err := userClient.GetUserList(context.Background(), &proto2.PageInfoRequest{
		Pn:    1,
		PSize: 3,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println()
	for _, user := range rsp.Data {

		fmt.Println(user.NickName, user.Id, user.Mobile, user.PassWord)
		checkRsp, err := userClient.CheckPassword(context.Background(), &proto2.PasswordCheckInfo{
			PassWord:          "admin123",
			EncryptedPassword: user.PassWord,
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(checkRsp.Success)

	}
}

//测试添加用户接口

func TestCreateUser(t *testing.T) {
	conn, err = grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		panic(err)
	}
	userClient = proto2.NewUserClient(conn)

	for i := 0; i < 10; i++ {
		rsp, err := userClient.CreateUser(context.Background(), &proto2.CreateUserInfo{
			NickName: fmt.Sprintf("keke%d", i),
			Mobile:   fmt.Sprintf("1882222333%d", i),
			PassWord: "xiaoke521",
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(rsp.Id)
	}
}
