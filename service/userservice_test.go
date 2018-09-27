package service_test

import (
	"context"
	"errors"
	"log"
	"net"
	"testing"
	"time"

	. "github.com/golang/mock/gomock"
	"github.com/grpc/grpc-go/test/bufconn"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/xjplke/grpctest/mock"
	pb "github.com/xjplke/grpctest/proto"
	"github.com/xjplke/grpctest/service"
	"google.golang.org/grpc"
)

//建议先看看这几篇文章
//https://www.jianshu.com/p/e3b2b1194830
//https://www.jianshu.com/p/70a93a9ed186
//https://www.jianshu.com/p/f4e773a1b11f
//https://www.jianshu.com/p/2f675d5e334e

const bufSize = 1024 * 1024

var lis *bufconn.Listener

var userService = &service.UserService{}

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()

	pb.RegisterUserServiceServer(s, userService)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(string, time.Duration) (net.Conn, error) {
	return lis.Dial()
}

func TestCreateUser(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := pb.NewUserServiceClient(conn)

	Convey("Create User", t, func() {
		ctrl := NewController(t)
		defer ctrl.Finish()
		mockRepo := mock.NewMockIUserRepo(ctrl)
		userService.Repo = mockRepo
		mockRepo.EXPECT().Save(Any()).Return(int64(123), nil)

		log.Println("Before Create User")
		resp, err := client.CreateUser(ctx, &pb.CreateUserRequest{Username: "test1", Nickname: "nickname1", Password: "123qwe"})
		So(err, ShouldEqual, nil)
		So(resp, ShouldNotEqual, nil)
		So(resp.Status, ShouldEqual, 0)

		Convey("Can't Create User with same Username", func() {
			mockRepo.EXPECT().Save(Any()).Return(int64(0), errors.New("duplicate Username"))

			_, err := client.CreateUser(ctx, &pb.CreateUserRequest{Username: "test1", Nickname: "nickname1", Password: "123qwe"})
			So(err, ShouldNotEqual, nil)
		})

		Convey("Update User Password", func() {
			mockRepo.EXPECT().ChangeNickname(Any(), Any()).Return(nil)

			_, err := client.ChangeNickname(ctx, &pb.ChangeNicknameRequest{Id: 123, Nickname: "nickname1"})
			So(err, ShouldEqual, nil)
		})
	})
}
