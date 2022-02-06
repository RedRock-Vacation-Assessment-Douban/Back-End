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
	celebrityIdString := ctx.Param("celebrity_id") //输入影人id
	celebrityId, err := strconv.Atoi(celebrityIdString)
	if err != nil {
		fmt.Println("celebrity id string to int err: ", err)
		tool.RespErrorWithDate(ctx, "celebrity_id格式有误")
		return
	}

	//根据celebrityId拿到celebrity
	celebrity, err := service.GetCelebrityById(celebrityId)
	if err != nil {
		fmt.Println("get celebrity by id err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessfulWithDate(ctx, celebrity)
}
