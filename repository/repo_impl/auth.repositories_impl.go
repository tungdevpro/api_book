package repoimpl

import (
	"api_book/models"
	"api_book/repository"
	"errors"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
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

func (repo *AuthRepositoriesImpl) Login(ctx *gin.Context) {}
func (repo *AuthRepositoriesImpl) Register(ctx *gin.Context, param models.User) (any, error) {
	col := repo.client.Database("bookz").Collection("users")
	filter := bson.D{{Key: "email", Value: param.Email}}

	var user models.User

	col.InsertOne(ctx, param)

	if err := col.FindOne(ctx, filter).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return "", errors.New("record does not exist")
		}
		return "", err
	}

	return "sucess", nil
}
func (repo *AuthRepositoriesImpl) ForgotPassword(ctx *gin.Context) {}
