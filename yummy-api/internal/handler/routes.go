package handler

import (
	db "yummy/internal/db/sqlc"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(engine *gin.Engine, queries *db.Queries) {
	router := engine.Group("/api/v1").Group("/foods")
	handler := &FoodHandler{queries: queries}
	router.GET("", handler.ListFoods)
	router.POST("", handler.CreateFoodItem)
	router.GET(":id", handler.GetFoodItem)
}
