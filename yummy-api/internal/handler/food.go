package handler

import (
	"database/sql"
	"path/filepath"
	"strconv"
	"strings"
	"yummy/internal/config"
	"yummy/internal/db/jet/yummy/public/model"
	"yummy/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FoodHandler struct {
	repo repository.FoodRepository
}

type FoodItem struct {
	ID        int32    `json:"id"`
	Name      string   `json:"name"`
	Caption   string   `json:"caption"`
	Rating    *float64 `json:"rating"`
	PhotoPath string   `json:"photo_path"`
}

func toFoodItem(f model.FoodItems) FoodItem {
	return FoodItem{
		ID:        f.ID,
		Name:      f.Name,
		Caption:   f.Caption,
		Rating:    f.Rating,
		PhotoPath: config.Config.BaseURL + f.PhotoPath,
	}
}

func (h *FoodHandler) ListFoods(c *gin.Context) {
	results, err := h.repo.List()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	res := make([]FoodItem, len(results))
	for i, f := range results {
		res[i] = toFoodItem(f)
	}

	c.JSON(200, gin.H{"data": res})
}

func (h *FoodHandler) GetFoodItem(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		return
	}

	result, err := h.repo.GetByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{"error": "not found"})
		} else {
			c.JSON(500, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(200, toFoodItem(result))
}

func (h *FoodHandler) CreateFoodItem(c *gin.Context) {
	for _, key := range []string{"name", "caption"} {
		if c.PostForm(key) == "" {
			c.JSON(400, gin.H{"error": key + " is required"})
			return
		}
	}

	rating, ok := parseRating(c)
	if !ok {
		return
	}

	photoPath, ok := savePhoto(c)
	if !ok {
		return
	}

	result, err := h.repo.Create(c.PostForm("name"), c.PostForm("caption"), photoPath, rating)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, toFoodItem(result))
}

func (h *FoodHandler) UpdateFoodItem(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		return
	}

	for _, key := range []string{"name", "caption"} {
		if c.PostForm(key) == "" {
			c.JSON(400, gin.H{"error": key + " is required"})
			return
		}
	}

	rating, ok := parseRating(c)
	if !ok {
		return
	}

	existing, err := h.repo.GetByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{"error": "not found"})
		} else {
			c.JSON(500, gin.H{"error": err.Error()})
		}
		return
	}

	photoPath := existing.PhotoPath
	if _, err := c.FormFile("photo"); err == nil {
		photoPath, ok = savePhoto(c)
		if !ok {
			return
		}
	}

	result, err := h.repo.Update(id, c.PostForm("name"), c.PostForm("caption"), photoPath, rating)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, toFoodItem(result))
}

func (h *FoodHandler) DeleteFoodItem(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		return
	}

	result, err := h.repo.Delete(id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{"error": "not found"})
		} else {
			c.JSON(500, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(200, toFoodItem(result))
}

func parseID(c *gin.Context) (int32, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
	}
	return int32(id), err
}

func parseRating(c *gin.Context) (*float64, bool) {
	r := c.PostForm("rating")
	if r == "" {
		return nil, true
	}
	val, err := strconv.ParseFloat(r, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid rating"})
		return nil, false
	}
	return &val, true
}

func savePhoto(c *gin.Context) (string, bool) {
	file, err := c.FormFile("photo")
	if err != nil {
		c.JSON(400, gin.H{"error": "photo is required"})
		return "", false
	}
	ext := filepath.Ext(file.Filename)
	filename := strings.TrimSuffix(file.Filename, ext)
	randomId, err := uuid.NewV7()
	if err != nil {
		c.JSON(500, gin.H{"error": "photo id creation failed"})
		return "", false
	}
	dest := "./uploads/" + filename + "-" + randomId.String() + ext
	path := "/uploads/" + filename + "-" + randomId.String() + ext
	if err = c.SaveUploadedFile(file, dest); err != nil {
		c.JSON(500, gin.H{"error": "photo upload failed"})
		return "", false
	}
	return path, true
}
