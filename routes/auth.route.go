package routes

import (
	"api_book/config"
	"api_book/handler"
	impl "api_book/repository/repo_impl"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func AuthRouter(engine *gin.Engine, group *gin.RouterGroup, client *mongo.Client) {
	authHandler := handler.AuthHandler{
		AuthRepository: impl.NewAuthRepository(client),
	}

	route := group.Group("/auth")
	{
		route.POST(config.Login, authHandler.Login)
		route.POST(config.Register, authHandler.Register)
	}
}
