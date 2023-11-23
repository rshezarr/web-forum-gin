package middleware

import (
	"forum/internal/model/error_model"
	"forum/internal/service/user_svc"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const userCtx = "user_id"

type Middlewarer interface {
	AuthMiddleware(c *gin.Context)
}

type middleware struct {
	svc user_svc.Userer
}

func NewMiddleware(svc user_svc.Userer) Middlewarer {
	return &middleware{
		svc: svc,
	}
}

func (m *middleware) AuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		error_model.NewErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		error_model.NewErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	if len(headerParts[1]) == 0 {
		error_model.NewErrorResponse(c, http.StatusUnauthorized, "token is empty")
		return
	}

	userID, err := m.svc.ParseToken(headerParts[1])
	if err != nil {
		error_model.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userID)
}
