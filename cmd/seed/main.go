package main

import (
	"context"
	"database/sql"
	"log"
	"yummy/internal/config"
	db "yummy/internal/db/sqlc"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var foodItems = []db.CreateFoodItemParams{
	{
		Name:      "Pizza",
		Caption:   "Delicious cheese pizza",
		Rating:    sql.NullFloat64{Float64: 4.5, Valid: true},
		PhotoPath: "uploads/images/pizza.png",
	},
	{
		Name:      "Burger",
		Caption:   "Juicy potato patty burger",
		Rating:    sql.NullFloat64{Float64: 4.2, Valid: true},
		PhotoPath: "uploads/images/burger.png",
	},
	{
		Name:      "Pasta",
		Caption:   "Creamy mix sauce pasta",
		Rating:    sql.NullFloat64{Float64: 4.8, Valid: true},
		PhotoPath: "uploads/images/pasta.png",
	},
}

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

	for _, item := range foodItems {
		_, err := queries.CreateFoodItem(context.Background(), item)
		if err != nil {
			log.Printf("Failed to insert food item %s: %v", item.Name, err)
		} else {
			log.Printf("Inserted food item: %s", item.Name)
		}
	}
}
