package service

import (
	"context"
	"errors"

	"fmt"

	"github.com/xjplke/grpctest/model"
	pb "github.com/xjplke/grpctest/proto"
)

//IUserRepo IUserRepo
type IUserRepo interface {
	Save(*model.User) (int64, error)
	ChangeNickname(id int64, nickname string) error
	CheckPassword(username, password string) error
}

//UserService UserService
type UserService struct {
	Repo IUserRepo
}

//CreateUser CreateUser
func (us *UserService) CreateUser(c context.Context, req *pb.CreateUserRequest) (*pb.UserReply, error) {
	um := &model.User{
		Username: req.Username,
		Password: req.Password,
		Nickname: req.Nickname,
	}
	fmt.Println("us.Repo = ", us.Repo)
	_, err := us.Repo.Save(um)

	resp := &pb.UserReply{
		Status: 0,
		Msgid:  "MSG_1",
	}
	return resp, err
}

//ChangeNickname ChangeNickname
func (us *UserService) ChangeNickname(c context.Context, req *pb.ChangeNicknameRequest) (*pb.UserReply, error) {
	err := us.Repo.ChangeNickname(req.Id, req.Nickname)
	rep := &pb.UserReply{}
	if err == nil {
		rep.Status = 0
	} else {
		rep.Status = 1
		rep.Msgid = "MSG_2"
	}
	return rep, err
}

//ChangePassword ChangePassword
func (us *UserService) ChangePassword(context.Context, *pb.ChangePasswordRequest) (*pb.UserReply, error) {
	return nil, errors.New("Unknow error")
}

//CheckPassword CheckPassword
func (us *UserService) CheckPassword(context.Context, *pb.CheckPasswordRequest) (*pb.UserReply, error) {
	return nil, errors.New("Unknow error")
}
