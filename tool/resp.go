package tool

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RespErrorWithDate(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"info": data,
	})
}

func RespInternalError(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"info": "服务器错误",
	})
}

func RespSuccessful(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"info": "成功",
	})
}

func RespSuccessfulWithDate(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"info": "成功",
		"data": data,
	})
}

func RespSuccessfulWithTwoDate(ctx *gin.Context, celebrity interface{}, movie interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"celebrity": celebrity,
		"movie":     movie,
	})
}

func RespSuccessfulWithTwoDate2(ctx *gin.Context, name interface{}, token interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"name":  name,
		"token": token,
	})
}
