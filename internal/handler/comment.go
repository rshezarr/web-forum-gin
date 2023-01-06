package handler

import (
	"errors"
	"forum/internal/model"
	"forum/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createComment(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	postId, err := strconv.Atoi(c.Param("post_id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "invalid id param")
		return
	}

	var comment model.Comment
	if err := c.BindJSON(&comment); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	comment.PostID = postId
	comment.UserID = userId

	id, err := h.service.Comment.Create(comment)
	if err != nil {
		if errors.Is(err, service.ErrInvalidComment) {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getCommentResponse struct {
	Data []model.Comment `json:"post_comment"`
}

func (h *Handler) getComment(c *gin.Context) {
	postId, err := strconv.Atoi(c.Param("post_id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "invalid id param")
		return
	}

	comments, err := h.service.Comment.GetByPostID(postId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getCommentResponse{
		Data: comments,
	})
}

func (h *Handler) updateComment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("comment_id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "invalid id param")
		return
	}

	var newComment model.Comment
	if err := c.BindJSON(&newComment); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	updId, err := h.service.Comment.Update(newComment.Content, id)
	if err != nil {
		if errors.Is(err, service.ErrInvalidComment) {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": updId,
	})
}

func (h *Handler) deleteComment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("comment_id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "invalid id param")
		return
	}

	if err := h.service.Comment.Delete(id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}
