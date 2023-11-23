package post_ctrl

import (
	"errors"
	"forum/internal/model"
	"forum/internal/model/error_model"
	"forum/internal/service/post_svc"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Poster interface {
	Posts(c *gin.Context)
	CreatePost(c *gin.Context)
	UpdatePost(c *gin.Context)
	DeletePost(c *gin.Context)
}

type postController struct {
	svc post_svc.Poster
}

func NewPostController(svc post_svc.Poster) Poster {
	return &postController{
		svc: svc,
	}
}

type getAllPostsResponse struct {
	Data []model.Post `json:"data"`
}

func (h *postController) Posts(c *gin.Context) {
	posts, err := h.svc.GetAll()
	if err != nil {
		error_model.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, getAllPostsResponse{
		Data: posts,
	})
}

func (h *postController) CreatePost(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		error_model.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var post model.Post
	if err := c.BindJSON(&post); err != nil {
		error_model.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	post.UserID = userId
	post.CreationTime = time.Now()

	id, err := h.svc.Create(post)
	if err != nil {
		if errors.Is(err, post_svc.ErrInvalidPostTitle) ||
			errors.Is(err, post_svc.ErrInvalidPostContent) ||
			errors.Is(err, post_svc.ErrPostTitleLen) ||
			errors.Is(err, post_svc.ErrPostContentLen) {
			error_model.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		error_model.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *postController) UpdatePost(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		error_model.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	postId, err := strconv.Atoi(c.Param("post_id"))
	if err != nil {
		error_model.NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	newPost := model.Post{
		ID:     postId,
		UserID: userId,
	}

	if err := c.BindJSON(&newPost); err != nil {
		error_model.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.svc.Update(newPost)
	if err != nil {
		error_model.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *postController) DeletePost(c *gin.Context) {
	postId, err := strconv.Atoi(c.Param("post_id"))
	if err != nil {
		error_model.NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	if err := h.svc.Delete(postId); err != nil {
		error_model.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}
