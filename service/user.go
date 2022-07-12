package service

import (
	"context"
	"fmt"
	user "helloKit/model"
)

type UserSvc interface {
	Create(ctx context.Context, req *user.CreateReq) (*user.CreateResp, error)
	Delete(ctx context.Context, req *user.DeleteReq) (*user.DeleteResp, error)
}

type userSvc struct {
}

func (u *userSvc) Create(ctx context.Context, req *user.CreateReq) (*user.CreateResp, error) {
	fmt.Println("Create")
	resp := user.NewCreateResp()
	resp.Code = 200
	resp.Msg = "success"
	resp.Data.Id = "12"
	resp.Data.Name = req.Name
	resp.Data.Age = req.Age
	return resp, nil
}

func (s *userSvc) Delete(context.Context, *user.DeleteReq) (*user.DeleteResp, error) {
	fmt.Println("Delete")
	return user.NewDeleteResp(), nil
}

func NewUserSvc() UserSvc {
	var svc = &userSvc{}
	{
		// middleware
	}
	return svc
}
