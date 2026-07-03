package handler

import (
	"database/sql"
	"path/filepath"
	"strconv"
	"strings"
	"yummy/internal/config"
	db "yummy/internal/db/sqlc"
	"yummy/internal/utils/nullable"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

func (handler *FoodHandler) ListFoods(c *gin.Context) {
	foods, err := handler.queries.ListFoods(c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if foods == nil {
		foods = []db.FoodItem{}
	}

	res := make([]FoodItem, len(foods))
	for i, food := range foods {
		res[i] = toFoodItem(food)
	}

	c.JSON(200, gin.H{"data": res})
}

func (handler *FoodHandler) GetFoodItem(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	food, err := handler.queries.GetFoodItem(c.Request.Context(), int32(id))
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{"error": "not found"})
		} else {
			c.JSON(500, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(200, toFoodItem(food))
}

func (handler *FoodHandler) CreateFoodItem(c *gin.Context) {
	for _, key := range []string{"name", "caption"} {
		if c.PostForm(key) == "" {
			c.JSON(400, gin.H{"error": key + " is required"})
			return
		}
	}

	var ratingPtr *float64
	if r := c.PostForm("rating"); r != "" {
		val, err := strconv.ParseFloat(r, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid rating"})
			return
		}
		ratingPtr = &val
	}

	file, err := c.FormFile("photo")
	if err != nil {
		c.JSON(400, gin.H{"error": "photo is required"})
		return
	}
	ext := filepath.Ext(file.Filename)
	filename := strings.TrimSuffix(file.Filename, ext)
	randomId, err := uuid.NewV7()
	if err != nil {
		c.JSON(500, gin.H{"error": "photo id creation failed"})
		return
	}
	destinationPath := "./uploads/" + filename + "-" + randomId.String() + ext
	photoPath := "/uploads/" + filename + "-" + randomId.String() + ext
	err = c.SaveUploadedFile(file, destinationPath)
	if err != nil {
		c.JSON(500, gin.H{"error": "photo upload failed"})
		return
	}

	payload := db.CreateFoodItemParams{
		Name:      c.PostForm("name"),
		Caption:   c.PostForm("caption"),
		Rating:    nullable.ToNullFloat64(ratingPtr),
		PhotoPath: photoPath,
	}

	food, err := handler.queries.CreateFoodItem(c.Request.Context(), payload)
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

	for _, key := range []string{"name", "caption"} {
		if c.PostForm(key) == "" {
			c.JSON(400, gin.H{"error": key + " is required"})
			return
		}
	}

	var ratingPtr *float64
	if r := c.PostForm("rating"); r != "" {
		val, err := strconv.ParseFloat(r, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid rating"})
			return
		}
		ratingPtr = &val
	}

	food, err := handler.queries.GetFoodItem(c.Request.Context(), int32(id))
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{"error": "not found"})
		} else {
			c.JSON(500, gin.H{"error": err.Error()})
		}
		return
	}
	photoPath := food.PhotoPath

	file, err := c.FormFile("photo")
	if err == nil {
		ext := filepath.Ext(file.Filename)
		filename := strings.TrimSuffix(file.Filename, ext)
		randomId, err := uuid.NewV7()
		if err != nil {
			c.JSON(500, gin.H{"error": "photo id creation failed"})
			return
		}
		destinationPath := "./uploads/" + filename + "-" + randomId.String() + ext
		photoPath = "/uploads/" + filename + "-" + randomId.String() + ext
		err = c.SaveUploadedFile(file, destinationPath)
		if err != nil {
			c.JSON(500, gin.H{"error": "photo upload failed"})
			return
		}
	}

	payload := db.UpdateFoodItemParams{
		ID:        int32(id),
		Name:      c.PostForm("name"),
		Caption:   c.PostForm("caption"),
		Rating:    nullable.ToNullFloat64(ratingPtr),
		PhotoPath: photoPath,
	}

	updatedFood, err := handler.queries.UpdateFoodItem(c.Request.Context(), payload)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, toFoodItem(updatedFood))
}

func (handler *FoodHandler) DeleteFoodItem(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	food, err := handler.queries.DeleteFoodItem(c.Request.Context(), int32(id))
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{"error": "not found"})
		} else {
			c.JSON(500, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(200, toFoodItem(food))
}
