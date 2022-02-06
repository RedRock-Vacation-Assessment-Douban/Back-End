package api

import (
	"douban/service"
	"douban/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func briefVideo(ctx *gin.Context) {
	videos, err := service.GetVideo1()
	if err != nil {
		fmt.Println("get err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessfulWithDate(ctx, videos)
}

func Video(ctx *gin.Context) {
	Sid := ctx.Param("recommend_id")
	id, _ := strconv.Atoi(Sid)
	videos, err := service.GetVideo2(id)
	if err != nil {
		fmt.Println("get err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessfulWithDate(ctx, videos)
}
