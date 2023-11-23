package v1

import (
	"forum/internal/http/v1/comment_ctrl"
	"forum/internal/http/v1/post_ctrl"
	"forum/internal/http/v1/user_ctrl"
	"forum/internal/middleware"
	"forum/internal/service"

	"github.com/gin-gonic/gin"
)

type Controllerer interface {
	InitRoutes() *gin.Engine
	StartRoutes()

	UserCtrlInit() user_ctrl.Userer
	PostCtrlInit() post_ctrl.Poster
	CommentCtrlInit() comment_ctrl.Commenter
}

type controller struct {
	service service.Servicer
	router  *gin.Engine
	middle  middleware.Middlewarer

	userCtrl    user_ctrl.Userer
	postCtrl    post_ctrl.Poster
	commentCtrl comment_ctrl.Commenter
}

func NewController(service service.Servicer, middle middleware.Middlewarer) Controllerer {
	return &controller{
		service: service,
		router:  gin.New(),
		middle:  middle,
	}
}

func (c *controller) InitRoutes() *gin.Engine {
	auth := c.router.Group("/user")
	{
		auth.POST("/sign-up", c.userCtrl.SignUp)
		auth.POST("/sign-in", c.userCtrl.SignIn)
	}

	api := c.router.Group("/api", c.middle.AuthMiddleware)
	{
		post := api.Group("/post")
		{
			post.GET("/", c.postCtrl.Posts)
			post.POST("/create", c.postCtrl.CreatePost)
			post.PUT("/update/:post_id", c.postCtrl.UpdatePost)
			post.DELETE("/delete/:post_id", c.postCtrl.DeletePost)
		}

		comment := api.Group("/comment")
		{
			comment.GET("/:post_id", c.commentCtrl.GetComment)
			comment.POST("/create/:post_id", c.commentCtrl.CreateComment)
			comment.PUT("/update/:comment_id", c.commentCtrl.UpdateComment)
			comment.DELETE("/delete/:comment_id", c.commentCtrl.DeleteComment)
		}
	}

	return c.router
}

func (c *controller) StartRoutes() {
	c.UserCtrlInit()
	c.PostCtrlInit()
	c.CommentCtrlInit()
}

func (c *controller) UserCtrlInit() user_ctrl.Userer {
	if c.userCtrl == nil {
		c.userCtrl = user_ctrl.NewUserController(c.service.UserSvcInit())
	}

	return c.userCtrl
}

func (c *controller) PostCtrlInit() post_ctrl.Poster {
	if c.postCtrl == nil {
		c.postCtrl = post_ctrl.NewPostController(c.service.PostSvcInit())
	}

	return c.postCtrl
}

func (c *controller) CommentCtrlInit() comment_ctrl.Commenter {
	if c.commentCtrl == nil {
		c.commentCtrl = comment_ctrl.NewCommentController(c.service.CommentSvcInit())
	}

	return c.commentCtrl
}
