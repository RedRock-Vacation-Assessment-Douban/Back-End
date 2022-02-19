package api

import (
	"database/sql"
	"douban/model"
	"douban/service"
	"douban/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

// filmCommentDetail 影评详细信息和其下属评论
func filmCommentDetail(ctx *gin.Context) {
	filmCommentIdString := ctx.Param("filmcomment_id") //输入影评id
	filmCommentId, err := strconv.Atoi(filmCommentIdString)
	if err != nil {
		fmt.Println("filmComment id string to int err: ", err)
		tool.RespErrorWithDate(ctx, "filmcomment_id格式有误")
		return
	}

	//根据filmCommentId拿到filmComment
	filmComment, err := service.GetFCById(filmCommentId)
	if err != nil {
		fmt.Println("get filmComment by id err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	//找到它的评论
	filmCommentReplys, err := service.GetFilmCommentReply(filmCommentId)
	if err != nil {
		if err != sql.ErrNoRows {
			fmt.Println("get filmComment comments err: ", err)
			tool.RespInternalError(ctx)
			return
		}
	}

	var filmCommentDetail model.FilmCommentDetail
	filmCommentDetail.FilmComment = filmComment
	filmCommentDetail.FilmCommentReplys = filmCommentReplys

	tool.RespSuccessfulWithDate(ctx, filmCommentDetail)
}

func briefFilmComments(ctx *gin.Context) {
	movieIdString := ctx.Param("movie_id")
	movieId, _ := strconv.Atoi(movieIdString)
	filmComments, err := service.GetFilmComments(movieId)
	if err != nil {
		fmt.Println("get filmComments err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessfulWithDate(ctx, filmComments)
}

// addFilmComment 添加影评
func addFilmComment(ctx *gin.Context) {
	iUsername, _ := ctx.Get("username") //得到用户名
	name := iUsername.(string)
	movieIdString := ctx.Param("movie_id")
	movieId, err := strconv.Atoi(movieIdString)
	context := ctx.PostForm("context")
	starNumString := ctx.PostForm("star_num")
	starNum, _ := strconv.Atoi(starNumString)
	MovieName, err := service.GetMNById(movieId)
	URL, err := service.GetURLByMId(movieId)
	filmComment := model.FilmComment{
		MovieId:   movieId,
		Context:   context,
		Name:      name,
		StarNum:   starNum,
		PostTime:  time.Now(),
		MovieName: MovieName,
		URL:       URL,
	}

	err = service.AddFilmComment(filmComment)
	if err != nil {
		fmt.Println("add filmComment err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessful(ctx)
}

// deleteFilmComment 删除影评
func deleteFilmComment(ctx *gin.Context) {
	filmCommentIdString := ctx.Param("filmcomment_id")
	filmCommentId, err := strconv.Atoi(filmCommentIdString)
	filmCommentNameString, _ := ctx.Get("username")
	nameString, _ := service.GetNameByFCId(filmCommentId)
	//必须用户名相同,无法删除他人话题
	if filmCommentNameString == nameString {
		if err != nil {
			fmt.Println("filmComment id string to int err: ", err)
			tool.RespErrorWithDate(ctx, "filmComment_id格式有误")
			return
		}
		err = service.DeleteFilmComment(filmCommentId)
		if err != nil {
			fmt.Println("delete filmComment err: ", err)
			tool.RespInternalError(ctx)
			return
		}
		tool.RespSuccessful(ctx)
	} else {
		tool.RespErrorWithDate(ctx, "不能删除他人评论")
	}
}

// filmCommentLikes 影评点赞
func filmCommentLikes(ctx *gin.Context) {
	filmCommentIdString := ctx.Param("filmcomment_id")
	filmCommentId, err := strconv.Atoi(filmCommentIdString)
	if err != nil {
		fmt.Println("filmComment id string to int err: ", err)
		tool.RespErrorWithDate(ctx, "filmComment_id格式有误")
		return
	}
	err = service.FilmCommentLikes(filmCommentId)
	if err != nil {
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessful(ctx)
}

// filmCommentDown 影评点踩
func filmCommentDown(ctx *gin.Context) {
	filmCommentIdString := ctx.Param("filmcomment_id")
	filmCommentId, err := strconv.Atoi(filmCommentIdString)
	if err != nil {
		fmt.Println("filmComment id string to int err: ", err)
		tool.RespErrorWithDate(ctx, "filmComment_id格式有误")
		return
	}
	err = service.FilmCommentDown(filmCommentId)
	if err != nil {
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessful(ctx)
}

func mostPopularFC(ctx *gin.Context) {
	MPFCs, err := service.GetMostPopular()
	if err != nil {
		fmt.Println("get MPFC err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessfulWithDate(ctx, MPFCs)
}

func briefFilmCommentsByUsername(ctx *gin.Context) {
	iUsername, _ := ctx.Get("username")
	username := iUsername.(string)
	filmComments, err := service.GetFilmCommentsByUsername(username)
	if err != nil {
		fmt.Println("get filmComments err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessfulWithDate(ctx, filmComments)
}
