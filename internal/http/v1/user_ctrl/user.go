package user_ctrl

import (
	"forum/internal/model"
	"forum/internal/model/error_model"
	"forum/internal/service/user_svc"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Userer interface {
	SignUp(c *gin.Context)
	SignIn(c *gin.Context)
}

type userController struct {
	svc user_svc.Userer
}

func NewUserController(svc user_svc.Userer) Userer {
	return &userController{
		svc: svc,
	}
}

func (h *userController) SignUp(c *gin.Context) {
	user := new(model.UserDto)

	if err := c.BindJSON(user); err != nil {
		error_model.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := h.svc.Create(user)
	if err != nil {
		error_model.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"user_id": id,
	})

}

func (h *userController) SignIn(c *gin.Context) {
	var user model.User

	if err := c.BindJSON(&user); err != nil {
		error_model.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	token, err := h.svc.GenerateToken(user.Email, user.Password)
	if err != nil {
		error_model.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
