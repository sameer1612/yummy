package main

import (
	"yummy/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	handler.RegisterRoutes(engine)
	engine.Run(":8000")
}
