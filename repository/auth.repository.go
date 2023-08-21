package repository

import "github.com/gin-gonic/gin"

type AuthRepository interface {
	Login(*gin.Context)
	Register(*gin.Context)
	ForgotPassword(*gin.Context)
}
