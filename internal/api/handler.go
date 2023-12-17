package api

import (
	"auth/internal/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Handler struct {
	srv    user.UserService
	logger *zap.Logger
}

func NewHandler(srv user.UserService, logger *zap.Logger) *Handler {
	return &Handler{srv: srv, logger: logger}
}

func (h *Handler) SignUp(ctx *gin.Context) {
	var user user.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "invalid JSON body",
		})
		return
	}
	newUser, err := h.srv.SignUp(ctx.Request.Context(), user.Username, user.Email, user.Password)
	if err != nil {
		h.logger.Error("error in signUp", zap.Error(err))
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "something went wrong",
		})
		return
	}
	ctx.JSON(http.StatusCreated, newUser)
}
