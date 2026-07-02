package main

import (
	"database/sql"
	"log"
	"yummy/internal/config"
	db "yummy/internal/db/sqlc"
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

	dbSQL, err := sql.Open("pgx", cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}

	if err := dbSQL.Ping(); err != nil {
		log.Fatal(err)
	}

	queries := db.New(dbSQL)

	engine := gin.Default()
	engine.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:4200"},
	}))
	handler.RegisterRoutes(engine, queries)
	engine.Run(cfg.Port)
}
