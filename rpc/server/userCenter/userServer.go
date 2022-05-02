package main

import (
	"context"
	"database/sql"
	"douban/api"
	"douban/dao"
	"douban/model"
	"douban/rpc/pb/userCenter"
	"douban/service/etcd"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type server struct {
	userCenter.UnimplementedRegisterCenterServer
}

func main() {
	listen, err := net.Listen("tcp", ":50001")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	userCenter.RegisterRegisterCenterServer(s, &server{})
	centerEtcd := etcd.NewCenterEtcd("127.0.0.1:2379")

	centerEtcd.AddNode(etcd.NewNode("douban", "server", "127.0.0.1:50001"))
	go func() {
		for {
			select {
			case <-centerEtcd.AliveChan:
				log.Println("续约")
			}
		}
	}()
	fmt.Println("正在监听......")
	if err = s.Serve(listen); err != nil {
		log.Println(err)
		return
	}
}

func (s *server) Register(ctx context.Context, req *userCenter.RegisterReq) (res *userCenter.RegisterRes, err error) {
	userInfo := model.User{
		Name:     req.Username,
		Password: req.Password,
		Question: req.Question,
		Answer:   req.Answer,
	}
	res = &userCenter.RegisterRes{}
	err = dao.Insert(userInfo)
	if err != nil {
		res.ID = 0
		res.Status = false
		return res, err
	}
	res.Status = true
	return res, nil
}

var f = true

func (s *server) Login(ctx context.Context, req *userCenter.LoginReq) (res *userCenter.LoginRes, err error) {
	res = &userCenter.LoginRes{}
	user, err := dao.SelectUserByUsername(req.Username)

	if err != nil {
		if err == sql.ErrNoRows {
			f = false
		}
		f = true
	}
	if user.Password != req.Password {
		f = false
	}
	f = true
	if err != nil {
		fmt.Println("judge password correct err: ", err)
		return
	}

	if !f {
		return
	}

	//jwt
	c := model.MyClaims{
		Username: req.Username,
		Password: req.Password,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 60,
			ExpiresAt: time.Now().Unix() + 6000000,
			Issuer:    "YuanXinHao",
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	s0, err := t.SignedString(api.MySigningKey)
	fmt.Println(s0)
	return res, err
}
