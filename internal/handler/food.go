package handler

import (
	db "yummy/internal/db/sqlc"

	"github.com/gin-gonic/gin"
)

type FoodHandler struct {
	q *db.Queries
}

func (h *FoodHandler) ListFoods(c *gin.Context) {
	foods, err := h.q.ListFoods(c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if foods == nil {
		foods = []db.FoodItem{}
	}

	c.JSON(200, gin.H{"data": foods})
}
