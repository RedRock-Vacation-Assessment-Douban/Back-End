package rpc

import (
	"context"
	"douban/rpc/pb/changePW"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func CallChangePw(endpoint string, token string, newPassword string) (res *changePW.ChangeRes) {
	conn, err := grpc.Dial(endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(err)
		return &changePW.ChangeRes{
			Status: false,
			Msg:    err.Error(),
		}
	}
	defer conn.Close()
	client := changePW.NewChangePWClient(conn)
	res, err = client.ChangePW(context.Background(),
		&changePW.ChangeReq{
			Token:       token,
			NewPassword: newPassword,
		},
	)
	if err != nil {
		log.Println(err)
		return &changePW.ChangeRes{
			Status: false,
			Msg:    err.Error(),
		}
	}
	fmt.Println("res = ", res)
	return res
}
