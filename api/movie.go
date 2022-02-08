package api

import (
	"douban/service"
	"douban/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// movieDetail
func movieDetail(ctx *gin.Context) {
	movieIdString := ctx.Param("movie_id") //输入电影id
	movieId, err := strconv.Atoi(movieIdString)
	if err != nil {
		fmt.Println("movie id string to int err: ", err)
		tool.RespErrorWithDate(ctx, "movie_id格式有误")
		return
	}

	//根据movieId拿到movie
	movie, err := service.GetMovieById(movieId)
	if err != nil {
		fmt.Println("get movie by id err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessfulWithDate(ctx, movie)
}

func briefMovies1(ctx *gin.Context) {
	movies, err := service.GetMovies1()
	if err != nil {
		fmt.Println("get movies err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessfulWithDate(ctx, movies)
}

func briefMovies2(ctx *gin.Context) {
	movies, err := service.GetMovies2()
	if err != nil {
		fmt.Println("get movies err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessfulWithDate(ctx, movies)
}

func briefMovies3(ctx *gin.Context) {
	movies, err := service.GetMovies3()
	if err != nil {
		fmt.Println("get movies err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessfulWithDate(ctx, movies)
}

func WTW(ctx *gin.Context) {
	iUsername, _ := ctx.Get("username") //得到用户名
	name := iUsername.(string)
	movieIdString := ctx.Param("movie_id") //输入电影id
	movieId, err := strconv.Atoi(movieIdString)
	URL := service.GetURLById(movieId)
	err = service.ChangeWTWId(name, movieId)
	if err != nil {
		fmt.Println("change  err: ", err)
		tool.RespInternalError(ctx)
		return
	}
	err2 := service.ChangeWTWURL(name, URL)
	if err2 != nil {
		fmt.Println("change  err: ", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessful(ctx)
}

func HW(ctx *gin.Context) {
	iUsername, _ := ctx.Get("username") //得到用户名
	name := iUsername.(string)
	movieIdString := ctx.Param("movie_id") //输入电影id
	movieId, err := strconv.Atoi(movieIdString)
	URL := service.GetURLById(movieId)
	err = service.ChangeHWId(name, movieId)
	if err != nil {
		fmt.Println("change  err: ", err)
		tool.RespInternalError(ctx)
		return
	}
	err2 := service.ChangeHWURL(name, URL)
	if err2 != nil {
		fmt.Println("change  err: ", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessful(ctx)
}

func rank1(ctx *gin.Context) {
	movie, err := service.GetMovie()
	if err != nil {
		fmt.Println("get movie err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessfulWithDate(ctx, movie)
}

func rank2(ctx *gin.Context) {
	movie, err := service.GetMovieRank1()
	if err != nil {
		fmt.Println("get movie err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessfulWithDate(ctx, movie)
}

func rank250(ctx *gin.Context) {
	movie, err := service.GetMovieRank250()
	if err != nil {
		fmt.Println("get movie err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessfulWithDate(ctx, movie)
}

func search(ctx *gin.Context) {
	SContext := ctx.PostForm("context")
	context := "%" + SContext + "%"
	movie, err := service.GetSearch(context)
	celebrity, err := service.GetSearch2(context)
	if err != nil {
		fmt.Println("get movie err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessfulWithTwoDate(ctx, celebrity, movie)
}

func rankUSA(ctx *gin.Context) {
	movie, err := service.GetMovieUSA()
	if err != nil {
		fmt.Println("get movie err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessfulWithDate(ctx, movie)
}

func classify(ctx *gin.Context) {
	_type := "%" + ctx.Param("type") + "%"
	_country := "%" + ctx.Param("country") + "%"
	movie, err := service.GetClassify(_type, _country)
	if err != nil {
		fmt.Println("get movie err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessfulWithDate(ctx, movie)
}

func all(ctx *gin.Context) {
	movie, err := service.GetAll()
	if err != nil {
		fmt.Println("get movie err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessfulWithDate(ctx, movie)
}

func classify1(ctx *gin.Context) {
	_country := "%" + ctx.Param("country") + "%"
	movie, err := service.GetClassify1(_country)
	if err != nil {
		fmt.Println("get movie err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessfulWithDate(ctx, movie)
}

func classify2(ctx *gin.Context) {
	_type := "%" + ctx.Param("type") + "%"
	movie, err := service.GetClassify2(_type)
	if err != nil {
		fmt.Println("get movie err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessfulWithDate(ctx, movie)
}
