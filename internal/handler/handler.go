package handler

import (
	"forum/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/user")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.authMiddleware)
	{
		post := api.Group("/post")
		{
			post.GET("/", h.posts)
			post.POST("/create", h.createPost)
			post.PUT("/update/:post_id", h.updatePost)
			post.DELETE("/delete/:post_id", h.deletePost)
			// post.POST("/like/:post_id", h.likePost)
			// post.POST("/dislike/:post_id", h.dislikePost)
		}
	}
	return router
}
