package v1

import (
	"forum/internal/service"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	service *service.Service
	handler *gin.Engine
}

func NewController(service *service.Service) *Controller {
	return &Controller{
		service: service,
		handler: gin.New(),
	}
}

func (h *Controller) InitRoutes() *gin.Engine {
	auth := h.handler.Group("/user")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := h.handler.Group("/api", h.authMiddleware)
	{
		post := api.Group("/post")
		{
			post.GET("/", h.posts)
			post.POST("/create", h.createPost)
			post.PUT("/update/:post_id", h.updatePost)
			post.DELETE("/delete/:post_id", h.deletePost)
		}

		comment := api.Group("/comment")
		{
			comment.GET("/:post_id", h.getComment)
			comment.POST("/create/:post_id", h.createComment)
			comment.PUT("/update/:comment_id", h.updateComment)
			comment.DELETE("/delete/:comment_id", h.deleteComment)
		}
	}

	return h.handler
}
