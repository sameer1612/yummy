package handler

import (
	db "yummy/internal/db/sqlc"

	"github.com/gin-gonic/gin"
)

type FoodHandler struct {
	queries *db.Queries
}

func (handler *FoodHandler) ListFoods(context *gin.Context) {
	foods, err := handler.queries.ListFoods(context.Request.Context())
	if err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if foods == nil {
		foods = []db.FoodItem{}
	}

	context.JSON(200, gin.H{"data": foods})
}
