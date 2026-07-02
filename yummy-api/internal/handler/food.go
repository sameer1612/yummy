package handler

import (
	"strconv"
	"yummy/internal/config"
	db "yummy/internal/db/sqlc"
	"yummy/internal/utils/nullable"

	"github.com/gin-gonic/gin"
)

type FoodHandler struct {
	queries *db.Queries
}

type FoodItem struct {
	ID        int32    `json:"id"`
	Name      string   `json:"name"`
	Caption   string   `json:"caption"`
	Rating    *float64 `json:"rating"`
	PhotoPath string   `json:"photo_path"`
}

type CreateFoodItemRequest struct {
	Name      string   `json:"name" binding:"required"`
	Caption   string   `json:"caption" binding:"required"`
	Rating    *float64 `json:"rating"`
	PhotoPath string   `json:"photo_path" binding:"required"`
}

type UpdateFoodItemRequest struct {
	Name      string   `json:"name" binding:"required"`
	Caption   string   `json:"caption" binding:"required"`
	Rating    *float64 `json:"rating"`
	PhotoPath string   `json:"photo_path" binding:"required"`
}

func toFoodItem(food db.FoodItem) FoodItem {
	return FoodItem{
		ID:        food.ID,
		Name:      food.Name,
		Caption:   food.Caption,
		Rating:    nullable.ToFloat64(food.Rating),
		PhotoPath: config.Config.BaseURL + food.PhotoPath,
	}
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
		res[i] = toFoodItem(food)
	}

	context.JSON(200, gin.H{"data": res})
}

func (handler *FoodHandler) GetFoodItem(context *gin.Context) {
	idParam := context.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		context.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	food, err := handler.queries.GetFoodItem(context.Request.Context(), int32(id))
	if err != nil {
		context.JSON(404, gin.H{"error": err.Error()})
		return
	}

	context.JSON(200, toFoodItem(food))
}

func (handler *FoodHandler) CreateFoodItem(c *gin.Context) {
	var req CreateFoodItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	food, err := handler.queries.CreateFoodItem(c.Request.Context(), db.CreateFoodItemParams{
		Name:      req.Name,
		Caption:   req.Caption,
		Rating:    nullable.ToNullFloat64(req.Rating),
		PhotoPath: req.PhotoPath,
	})
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, toFoodItem(food))
}

func (handler *FoodHandler) UpdateFoodItem(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	var req UpdateFoodItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	food, err := handler.queries.UpdateFoodItem(c.Request.Context(), db.UpdateFoodItemParams{
		ID:        int32(id),
		Name:      req.Name,
		Caption:   req.Caption,
		Rating:    nullable.ToNullFloat64(req.Rating),
		PhotoPath: req.PhotoPath,
	})

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, toFoodItem(food))
}
