package main

import (
	"log"
	"yummy/internal/config"
	"yummy/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	engine := gin.Default()
	handler.RegisterRoutes(engine)
	engine.Run(cfg.Port)
}
