package handler

import (
	"database/sql"
	"yummy/internal/repository"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(engine *gin.Engine, db *sql.DB) {
	router := engine.Group("/api/v1").Group("/foods")
	handler := &FoodHandler{repo: repository.NewRepository(db)}
	router.GET("", handler.ListFoods)
	router.POST("", handler.CreateFoodItem)
	router.GET(":id", handler.GetFoodItem)
	router.PUT(":id", handler.UpdateFoodItem)
	router.DELETE(":id", handler.DeleteFoodItem)
}
