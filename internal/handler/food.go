package handler

import (
	db "yummy/internal/db/sqlc"
	"yummy/internal/utils/nullable"

	"github.com/gin-gonic/gin"
)

type FoodItem struct {
	ID      int32    `json:"id"`
	Name    string   `json:"name"`
	Caption string   `json:"caption"`
	Rating  *float64 `json:"rating"`
}

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

	res := make([]FoodItem, len(foods))
	for i, food := range foods {
		res[i] = FoodItem{
			ID:      food.ID,
			Name:    food.Name,
			Caption: food.Caption,
			Rating:  nullable.NullableFloat(food.Rating),
		}
	}

	context.JSON(200, gin.H{"data": res})
}
