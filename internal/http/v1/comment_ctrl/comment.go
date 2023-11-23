package comment_ctrl

import (
	"errors"
	"forum/internal/model"
	"forum/internal/model/error_model"
	"forum/internal/service/comment_svc"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Commenter interface {
	CreateComment(c *gin.Context)
	GetComment(c *gin.Context)
	UpdateComment(c *gin.Context)
	DeleteComment(c *gin.Context)
}

type commentController struct {
	svc comment_svc.Commenter
}

func NewCommentController(svc comment_svc.Commenter) Commenter {
	return &commentController{
		svc: svc,
	}
}

func (h *commentController) CreateComment(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		error_model.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	postId, err := strconv.Atoi(c.Param("post_id"))
	if err != nil {
		error_model.NewErrorResponse(c, http.StatusInternalServerError, "invalid id param")
		return
	}

	var comment model.Comment
	if err := c.BindJSON(&comment); err != nil {
		error_model.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	comment.PostID = postId
	comment.UserID = userId

	id, err := h.svc.Create(comment)
	if err != nil {
		if errors.Is(err, comment_svc.ErrInvalidComment) {
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

type getCommentResponse struct {
	Data []model.Comment `json:"post_comment"`
}

func (h *commentController) GetComment(c *gin.Context) {
	postId, err := strconv.Atoi(c.Param("post_id"))
	if err != nil {
		error_model.NewErrorResponse(c, http.StatusInternalServerError, "invalid id param")
		return
	}

	comments, err := h.svc.GetByPostID(postId)
	if err != nil {
		error_model.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getCommentResponse{
		Data: comments,
	})
}

func (h *commentController) UpdateComment(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		error_model.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("comment_id"))
	if err != nil {
		error_model.NewErrorResponse(c, http.StatusInternalServerError, "invalid id param")
		return
	}

	var newComment model.Comment
	if err := c.BindJSON(&newComment); err != nil {
		error_model.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newComment.UserID = userId
	updId, err := h.svc.Update(newComment, id)
	if err != nil {
		if errors.Is(err, comment_svc.ErrInvalidComment) {
			error_model.NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		error_model.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": updId,
	})
}

func (h *commentController) DeleteComment(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		error_model.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("comment_id"))
	if err != nil {
		error_model.NewErrorResponse(c, http.StatusInternalServerError, "invalid id param")
		return
	}

	if err := h.svc.Delete(id, userId); err != nil {
		error_model.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"info": "content has ben deleted",
	})
}
