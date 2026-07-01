package handler

import (
	db "yummy/internal/db/sqlc"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(engine *gin.Engine, q *db.Queries) {
	router := engine.Group("/api/v1").Group("/foods")
	h := &FoodHandler{q: q}
	router.GET("", h.ListFoods)
}
