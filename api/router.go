package api

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func InitEngine() {
	engine := gin.Default()
	engine.Use(CORS())

	engine.Use(static.Serve("/", static.LocalFile("./static", false)))
	engine.POST("/register", register)       //注册
	engine.POST("/login", login)             //登陆
	engine.POST("/mibao", mibao)             //密保
	engine.POST("/mibao/question", question) //查询密保问题

	engine.GET("/brief1", briefMovies1) //电影页
	engine.GET("/brief2", briefMovies2)
	engine.GET("/brief3", briefMovies3)
	engine.GET("/recommend", briefVideo)      //推荐页
	engine.GET("/mostpopular", mostPopularFC) //最受欢迎影评

	engine.GET("/rank1", rank1)   //排行榜
	engine.GET("/rank2", rank2)   //排行榜
	engine.GET("/rank3", rank250) //排行榜
	engine.GET("/rank4", rankUSA) //排行榜

	engine.POST("/search", search) //搜索电影
	engine.POST("/classify", all)
	engine.POST("/classify/:type/:country", classify)
	engine.POST("/classify1/:country", classify1)
	engine.POST("/classify2/:type", classify2)
	engine.POST("/classifyrank/:type", classifyRank)

	movieGroup := engine.Group("/movie")
	movieGroup.Use(CORS())
	{
		movieGroup.GET("/:movie_id", movieDetail) //电影页
		movieGroup.GET("/:movie_id/:celebrity_id", celebrityDetail)
		{
			movieGroup.Use(JWTAuth)
			movieGroup.GET("/wtw/:movie_id", WTW) //想看
			movieGroup.GET("/hw/:movie_id", HW)   //看过
		}
	}

	recommendGroup := engine.Group("/recommend")
	recommendGroup.Use(CORS())
	{
		recommendGroup.GET("/:recommend_id", Video)
	}

	userGroup := engine.Group("/user")
	userGroup.Use(CORS())
	{
		userGroup.Use(JWTAuth)                      //需要token
		userGroup.POST("/password", changePassword) //修改密码
		userGroup.POST("/introduction", changeSI)   //修改自我介绍
		userGroup.GET("/user1", user)               //查看个人页面
		userGroup.GET("/user2", userMovie)
		userGroup.GET("/user3", briefFilmCommentsByUsername)
		userGroup.GET("/user4", briefShortCommentByUsername)
	}

	topicGroup := engine.Group("/topic")
	topicGroup.Use(CORS())
	{
		topicGroup.GET("/movie/:movie_id", briefTopics) //查看一部电影全部话题概略
		topicGroup.GET("/:topic_id", topicDetail)       //查看一条话题详细信息和其下属评论
		{
			topicGroup.Use(JWTAuth)                        //需要token
			topicGroup.POST("/:movie_id", addTopic)        //发布新话题
			topicGroup.DELETE("/:topic_id", deleteTopic)   //删除话题
			topicGroup.GET("/likes/:topic_id", topicLikes) //给话题点赞
		}
	}

	commentGroup := engine.Group("/comment")
	commentGroup.Use(CORS())
	{
		commentGroup.POST("/anonymity/:topic_id", addCommentAnonymity) //匿名评论
		{
			commentGroup.Use(JWTAuth)                            //需要token
			commentGroup.POST("/:topic_id", addComment)          //发送评论
			commentGroup.DELETE("/:comment_id", deleteComment)   //删除评论
			commentGroup.GET("/likes/:comment_id", commentLikes) //给评论点赞
		}
	}

	shortCommentGroup := engine.Group("/shortcomment")
	shortCommentGroup.Use(CORS())
	{
		shortCommentGroup.GET("/movie/:movie_id", briefShortComment) //查看一部电影全部短评
		{
			shortCommentGroup.Use(JWTAuth)                                      //需要token
			shortCommentGroup.POST("/:movie_id", addShortComment)               //发布新短评
			shortCommentGroup.DELETE("/:shortcomment_id", deleteShortComment)   //删除短评
			shortCommentGroup.GET("/likes/:shortcomment_id", shortCommentLikes) //给短评点赞
		}
	}

	filmCommentGroup := engine.Group("/filmcomment")
	filmCommentGroup.Use(CORS())
	{
		filmCommentGroup.GET("/movie/:movie_id", briefFilmComments) //查看一部电影全部影评概略
		filmCommentGroup.GET("/:filmcomment_id", filmCommentDetail) //查看一条影评详细信息和其下属评论
		{
			filmCommentGroup.Use(JWTAuth)                                    //需要token
			filmCommentGroup.POST("/:movie_id", addFilmComment)              //发布新影评
			filmCommentGroup.DELETE("/:filmcomment_id", deleteFilmComment)   //删除影评
			filmCommentGroup.GET("/likes/:filmcomment_id", filmCommentLikes) //给影评点赞
			filmCommentGroup.GET("/down/:filmcomment_id", filmCommentDown)   //给影评点踩
		}
	}

	filmCommentReplyGroup := engine.Group("/filmcomment_reply")
	filmCommentReplyGroup.Use(CORS())
	{
		filmCommentReplyGroup.POST("/anonymity/:filmcomment_id", addFilmCommentReplyAnonymity) //匿名评论
		{
			filmCommentReplyGroup.Use(JWTAuth)                                             //需要token
			filmCommentReplyGroup.POST("/:filmcomment_id", addFilmCommentReply)            //发送评论
			filmCommentReplyGroup.DELETE("/:filmcomment_reply_id", deleteFilmCommentReply) //删除评论
		}
	}

	//engine.Use(TlsHandler(8081))
	//err := engine.RunTLS(":8081", "/data/42.192.155.29_chain.crt", "/data/42.192.155.29_key.key")
	//if err != nil {
	//	return
	//}
	err := engine.Run(":8080")
	if err != nil {
		return
	}
}
