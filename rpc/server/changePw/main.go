package main

import (
	"context"
	"douban/api"
	"douban/dao"
	"douban/rpc/pb/changePW"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	changePW.UnimplementedChangePWServer
}

func main() {
	listen, err := net.Listen("tcp", ":50002")
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	s := grpc.NewServer()
	changePW.RegisterChangePWServer(s, &server{}) //注册这一个服务

	if err = s.Serve(listen); err != nil { //监听
		fmt.Println("err:", err)
		return
	}
}

func (s *server) ChangePW(ctx context.Context, req *changePW.ChangeReq) (res *changePW.ChangeRes, err error) {
	username := api.ParseToken(req.Token)
	res = &changePW.ChangeRes{
		Status: false,
		Msg:    "",
	}
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	if err = dao.UpdatePassword(username, req.NewPassword); err != nil {
		fmt.Println("err:", err)
		return
	}
	return
}
