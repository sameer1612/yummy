package handler

import (
	"database/sql"
	"path/filepath"
	"strconv"
	"strings"
	"yummy/internal/config"
	"yummy/internal/db/jet/yummy/public/model"
	"yummy/internal/db/jet/yummy/public/table"

	. "github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

type FoodHandler struct {
	db *sql.DB
}

type FoodItem struct {
	ID        int32    `json:"id"`
	Name      string   `json:"name"`
	Caption   string   `json:"caption"`
	Rating    *float64 `json:"rating"`
	PhotoPath string   `json:"photo_path"`
}

func toFoodItem(food model.FoodItems) FoodItem {
	return FoodItem{
		ID:        food.ID,
		Name:      food.Name,
		Caption:   food.Caption,
		Rating:    food.Rating,
		PhotoPath: config.Config.BaseURL + food.PhotoPath,
	}
}

var tbl = table.FoodItems

func (handler *FoodHandler) ListFoods(c *gin.Context) {
	var results []model.FoodItems
	err := SELECT(tbl.AllColumns).FROM(tbl).Query(handler.db, &results)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	res := make([]FoodItem, len(results))
	for i, food := range results {
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

	var result model.FoodItems
	err = SELECT(tbl.AllColumns).FROM(tbl).WHERE(tbl.ID.EQ(Int32(int32(id)))).Query(handler.db, &result)
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

	var result model.FoodItems
	err = tbl.INSERT(tbl.Name, tbl.Caption, tbl.Rating, tbl.PhotoPath).MODEL(model.FoodItems{
		Name:      c.PostForm("name"),
		Caption:   c.PostForm("caption"),
		Rating:    ratingPtr,
		PhotoPath: photoPath,
	}).RETURNING(tbl.AllColumns).Query(handler.db, &result)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, toFoodItem(result))
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

	var food model.FoodItems
	err = SELECT(tbl.AllColumns).FROM(tbl).WHERE(tbl.ID.EQ(Int32(int32(id)))).Query(handler.db, &food)
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

	var result model.FoodItems
	err = tbl.UPDATE(tbl.Name, tbl.Caption, tbl.Rating, tbl.PhotoPath).MODEL(model.FoodItems{
		Name:      c.PostForm("name"),
		Caption:   c.PostForm("caption"),
		Rating:    ratingPtr,
		PhotoPath: photoPath,
	}).WHERE(tbl.ID.EQ(Int32(int32(id)))).RETURNING(tbl.AllColumns).Query(handler.db, &result)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, toFoodItem(result))
}

func (handler *FoodHandler) DeleteFoodItem(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	var result model.FoodItems
	err = tbl.DELETE().WHERE(tbl.ID.EQ(Int32(int32(id)))).RETURNING(tbl.AllColumns).Query(handler.db, &result)
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
