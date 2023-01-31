package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vladmsnk/billing/internal/dto"
	"vladmsnk/billing/internal/usecase"
	"vladmsnk/billing/pkg/logger"
)

type AuthRoutes struct {
	t usecase.Auth
	l logger.Interface
}

func newAuthRoutes(handler *gin.RouterGroup, a usecase.Auth, l logger.Interface) {
	r := &AuthRoutes{a, l}

	handler.POST("/token", r.getToken)
	handler.POST("/user/register", r.createUser)

}

func (a *AuthRoutes) getToken(ctx *gin.Context) {

	var request dto.TokenRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	token, err := a.t.GenerateToken(ctx, request)
	if err != nil {
		//change
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})

}

func (a *AuthRoutes) createUser(ctx *gin.Context) {
	var user dto.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	err := a.t.CreateUser(ctx, user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
}
