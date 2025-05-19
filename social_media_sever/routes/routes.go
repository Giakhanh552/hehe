package routes

import (
	controllers "social_media_sever/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Post routes
	postCtl := controllers.NewPostController()
	r.POST("/posts", postCtl.CreatePost)
	r.GET("/posts", postCtl.GetPosts)
	r.GET("/posts/:id", postCtl.GetPostById)
	r.PUT("/posts/:id", postCtl.UpdatePost)
	r.DELETE("/posts/:id", postCtl.DeletePost)

	// Comment routes
	commentCtl := controllers.NewCommentController()
	r.POST("/comments", commentCtl.CreateComment)
	r.GET("/comments", commentCtl.GetComments)
	r.GET("/posts/:id/comments", commentCtl.GetCommentsByPostId)
	r.PUT("/comments/:id", commentCtl.UpdateComment)
	r.DELETE("/comments/:id", commentCtl.DeleteComment)

	return r
}
