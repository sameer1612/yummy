package handler

import (
	"yummy/internal/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(engine *gin.Engine, db *gorm.DB) {
	router := engine.Group("/api/v1").Group("/foods")
	handler := &FoodHandler{repo: repository.NewFoodRepository(db)}
	router.GET("", handler.ListFoods)
	router.POST("", handler.CreateFoodItem)
	router.GET(":id", handler.GetFoodItem)
	router.PUT(":id", handler.UpdateFoodItem)
	router.DELETE(":id", handler.DeleteFoodItem)
}
