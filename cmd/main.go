package main

import (
	"api_book/config"
	"api_book/db"
	"api_book/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Load()

	database := db.MongoDatabase{}
	client := database.RunConnect()
	database.Setup(client)

	r := gin.Default()
	v1 := r.Group("/v1")
	{
		routes.AuthRouter(r, v1, database.Client)
		routes.UserRouter(r, v1)
	}

	r.Run(":3002")
}
