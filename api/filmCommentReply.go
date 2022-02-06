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

// addFilmCommentReply 添加评论
func addFilmCommentReply(ctx *gin.Context) {
	iUsername, _ := ctx.Get("username")
	Name := iUsername.(string)

	context := ctx.PostForm("context")
	filmCommentIdString := ctx.Param("filmcomment_id") //评论的话题id
	FilmCommentId, err := strconv.Atoi(filmCommentIdString)
	if err != nil {
		fmt.Println("FilmComment id string to int err: ", err)
		tool.RespErrorWithDate(ctx, "影评id有误")
		return
	}

	FilmCommentReply := model.FilmCommentReply{
		FilmCommentId: FilmCommentId,
		Context:       context,
		Name:          Name,
		CommentTime:   time.Now(),
	}
	err = service.AddFilmCommentReply(FilmCommentReply)
	if err != nil {
		fmt.Println("add FilmCommentReply err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessful(ctx)
}

// addFilmCommentReplyAnonymity 匿名评论
func addFilmCommentReplyAnonymity(ctx *gin.Context) {
	Name := "Anonymity" //匿名评论用户名统一为Anonymity
	context := ctx.PostForm("context")
	filmCommentIdString := ctx.Param("filmcomment_id") //评论的话题id
	FilmCommentId, err := strconv.Atoi(filmCommentIdString)
	if err != nil {
		fmt.Println("FilmComment id string to int err: ", err)
		tool.RespErrorWithDate(ctx, "影评id有误")
		return
	}

	FilmCommentReply := model.FilmCommentReply{
		FilmCommentId: FilmCommentId,
		Context:       context,
		Name:          Name,
		CommentTime:   time.Now(),
	}
	err = service.AddFilmCommentReply(FilmCommentReply)
	if err != nil {
		fmt.Println("add FilmCommentReply err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessful(ctx)
}

// deleteFilmCommentReply 删除评论
func deleteFilmCommentReply(ctx *gin.Context) {
	filmCommentReplyIdString := ctx.Param("filmcomment_reply_id") //输入评论id
	filmCommentReplyId, err := strconv.Atoi(filmCommentReplyIdString)
	filmCommentReplyNameString, _ := ctx.Get("username") //取用户名
	nameString, _ := service.GetNameByFCRId(filmCommentReplyId)
	//不能删除他人的评论,将用户名进行判断
	if filmCommentReplyNameString == nameString {
		if err != nil {
			fmt.Println("FilmCommentReply id string to int err: ", err)
			tool.RespErrorWithDate(ctx, "FilmCommentReply_id格式有误")
			return
		}
		err = service.DeleteFilmCommentReply(filmCommentReplyId)
		if err != nil {
			fmt.Println("delete FilmCommentReply err: ", err)
			tool.RespInternalError(ctx)
			return
		}
		tool.RespSuccessful(ctx)
	} else {
		tool.RespErrorWithDate(ctx, "不能删除他人评论")
	}
}
