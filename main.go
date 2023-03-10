package main

import (
	"api/app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.Init(router)
	router.Run(":3000")
}
