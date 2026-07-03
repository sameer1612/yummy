package main

import (
	"database/sql"
	"log"
	"yummy/internal/config"
	"yummy/internal/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	config.Config = cfg

	dbSQL, err := sql.Open("pgx", cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}

	if err := dbSQL.Ping(); err != nil {
		log.Fatal(err)
	}

	engine := gin.Default()
	engine.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:4200"},
		AllowHeaders: []string{"*"},
		AllowMethods: []string{"*"},
	}))

	engine.Static("/uploads", "./uploads")

	handler.RegisterRoutes(engine, dbSQL)
	engine.Run(cfg.Port)
}
