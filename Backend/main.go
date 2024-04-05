package main

import (
	"backend/config"
	"backend/router"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig(".env")

	r := gin.Default()
	r.Use(cors.Default())
	api := r.Group("/v1")
	router.AuthRouter(api)
	r.Run(fmt.Sprintf(":%v", 8000))
}
