package api

import (
	"douban/model"
	"douban/service"
	"douban/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func briefShortComment(ctx *gin.Context) {
	movieIdString := ctx.Param("movie_id")
	movieId, _ := strconv.Atoi(movieIdString)
	shortComments, err := service.GetShortComment(movieId)
	if err != nil {
		fmt.Println("get shortComments err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessfulWithDate(ctx, shortComments)
}

func briefShortCommentByUsername(ctx *gin.Context) {
	iUsername, _ := ctx.Get("username")
	username := iUsername.(string)
	shortComments, err := service.GetShortCommentByUsername(username)
	if err != nil {
		fmt.Println("get shortComments err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessfulWithDate(ctx, shortComments)
}

// addShortComment 添加短评
func addShortComment(ctx *gin.Context) {
	iUsername, _ := ctx.Get("username") //得到用户名
	name := iUsername.(string)
	movieIdString := ctx.Param("movie_id")
	movieId, err := strconv.Atoi(movieIdString)
	context := ctx.PostForm("context")
	starNumString := ctx.PostForm("starNum")
	starNum, _ := strconv.Atoi(starNumString)
	status := ctx.PostForm("status")
	shortComment := model.ShortComment{
		MovieId:     movieId,
		Name:        name,
		Status:      status,
		StarNum:     starNum,
		CommentTime: time.Now(),
		Context:     context,
	}

	err = service.AddShortComment(shortComment)
	if err != nil {
		fmt.Println("add shortComment err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessful(ctx)
}

// deleteShortComment 删除短评
func deleteShortComment(ctx *gin.Context) {
	shortCommentIdString := ctx.Param("shortcomment_id")
	shortCommentId, err := strconv.Atoi(shortCommentIdString)
	shortCommentNameString, _ := ctx.Get("username")
	nameString, _ := service.GetNameBySCId(shortCommentId)
	//必须用户名相同,无法删除他人话题
	if shortCommentNameString == nameString {
		if err != nil {
			fmt.Println("shortComment id string to int err: ", err)
			tool.RespErrorWithDate(ctx, "shortcomment_id格式有误")
			return
		}
		err = service.DeleteShortComment(shortCommentId)
		if err != nil {
			fmt.Println("delete shortComment err: ", err)
			tool.RespInternalError(ctx)
			return
		}
		tool.RespSuccessful(ctx)
	} else {
		tool.RespErrorWithDate(ctx, "不能删除他人短评")
	}
}

// shortCommentLikes 话题点赞
func shortCommentLikes(ctx *gin.Context) {
	shortCommentIdString := ctx.Param("shortcomment_id")
	shortCommentId, err := strconv.Atoi(shortCommentIdString)
	if err != nil {
		fmt.Println("shortComment id string to int err: ", err)
		tool.RespErrorWithDate(ctx, "shortcomment_id格式有误")
		return
	}
	err = service.ShortCommentLikes(shortCommentId)
	if err != nil {
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessful(ctx)
}
