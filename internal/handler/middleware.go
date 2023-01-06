package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const userCtx = "user_id"

func (h *Handler) authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
			return
		}

		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
			return
		}

		if len(headerParts[1]) == 0 {
			newErrorResponse(c, http.StatusUnauthorized, "token is empty")
			return
		}

		userID, err := h.service.ParseToken(headerParts[1])
		if err != nil {
			newErrorResponse(c, http.StatusUnauthorized, err.Error())
			return
		}

		c.Set(userCtx, userID)
	}
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("user id is of invalid type")
	}

	return idInt, nil
}
