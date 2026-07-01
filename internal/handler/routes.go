package handler

import "github.com/gin-gonic/gin"

func RegisterRoutes(engine *gin.Engine) {
	router := engine.Group("/api/v1").Group("/foods")
	{
		router.GET("", ListFoods)
	}
}
