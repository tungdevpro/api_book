package repoimpl

import (
	"api_book/repository"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthRepositoriesImpl struct {
	client *mongo.Client
}

func NewAuthRepository(client *mongo.Client) repository.AuthRepository {
	return &AuthRepositoriesImpl{
		client: client,
	}
}

func (repo *AuthRepositoriesImpl) Login(ctx *gin.Context)          {}
func (repo *AuthRepositoriesImpl) Register(ctx *gin.Context)       {}
func (repo *AuthRepositoriesImpl) ForgotPassword(ctx *gin.Context) {}
