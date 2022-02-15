package api

import (
	"douban/service"
	"douban/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// celebrityDetail
func celebrityDetail(ctx *gin.Context) {
	movieIdString := ctx.Param("movie_id") //输入电影id
	movieId, err := strconv.Atoi(movieIdString)
	if err != nil {
		fmt.Println("movie id string to int err: ", err)
		tool.RespErrorWithDate(ctx, "movie_id格式有误")
		return
	}

	celebrityIdString := ctx.Param("celebrity_id") //输入影人id
	celebrityId, err := strconv.Atoi(celebrityIdString)
	if err != nil {
		fmt.Println("celebrity id string to int err: ", err)
		tool.RespErrorWithDate(ctx, "celebrity_id格式有误")
		return
	}

	//根据celebrityId拿到celebrity
	celebrity, err := service.GetCelebrityById(celebrityId, movieId)
	if err != nil {
		fmt.Println("get celebrity by id err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessfulWithDate(ctx, celebrity)
}
