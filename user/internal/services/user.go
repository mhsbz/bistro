package services

import (
	"context"
	"fmt"
	v1 "github.com/mhsbz/uinterface/user/v1"
)

type UserService struct {
	v1.UnimplementedUserServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (u *UserService) CreateUser(ctx context.Context, param *v1.UserStructInput) (*v1.UserStructOutput, error) {
	fmt.Println("i am here...")
	return &v1.UserStructOutput{User: &v1.UserStruct{
		Id:       1,
		UserName: "zhangsan",
	}}, nil
}
