package main

import (
	"api_book/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		routes.UserRouter(r, v1)
	}

	r.Run(":3001")
}
