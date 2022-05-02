package rpc

import (
	"context"
	"douban/model"
	"douban/rpc/pb/userCenter"
	"douban/service/etcd"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

// userLoginAndRegister
type userLR struct {
	userName string
	passWord string
	node     string
	conn     *grpc.ClientConn
}

func NewUserLrCtl(endpoint, projectName, chileNode string) (u *userLR) {
	w := etcd.NewWatcher(endpoint, projectName)
	node, err2 := w.MatchChildNode(chileNode)
	if err2 != nil {
		log.Println(err2)
		return
	}
	conn, err := grpc.Dial(node, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	userLrCtl := userLR{
		node: node,
		conn: conn,
	}
	return &userLrCtl
}

func (u *userLR) CallRegister(info model.User) *userCenter.RegisterRes {
	u.userName = info.Name
	u.passWord = info.Password
	client := userCenter.NewRegisterCenterClient(u.conn)
	return u.register(client)
}

func (u *userLR) CallLogin(info model.User) *userCenter.LoginRes {
	u.userName = info.Name
	u.passWord = info.Password
	client := userCenter.NewRegisterCenterClient(u.conn)
	return u.login(client)
}

func (user *userLR) register(client userCenter.RegisterCenterClient) *userCenter.RegisterRes {
	res, err := client.Register(context.Background(), &userCenter.RegisterReq{
		Username: user.userName,
		Password: user.passWord,
	})
	if err != nil {
		log.Println(err)
		return &userCenter.RegisterRes{}
	}
	if res.Status == true {
		fmt.Println("SUCCESS")
		return res
	}
	return &userCenter.RegisterRes{}
}

func (user *userLR) login(client userCenter.RegisterCenterClient) *userCenter.LoginRes {
	res, err := client.Login(context.Background(), &userCenter.LoginReq{
		Username: user.userName,
		Password: user.passWord,
	})
	if err != nil {
		log.Println(err)
		return &userCenter.LoginRes{
			Status: false,
			Token:  "",
		}
	}
	return res
}
