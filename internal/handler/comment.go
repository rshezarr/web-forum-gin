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

func (h *Handler) getComment(c *gin.Context) {

}

func (h *Handler) updateComment(c *gin.Context) {

}

func (h *Handler) deleteComment(c *gin.Context) {

}
