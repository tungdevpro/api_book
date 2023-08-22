package repository

import (
	"api_book/models"

	"github.com/gin-gonic/gin"
)

type AuthRepository interface {
	Login(*gin.Context)
	Register(*gin.Context, models.User) (any, error)
	ForgotPassword(*gin.Context)
}
