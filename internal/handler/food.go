package handler

import "github.com/gin-gonic/gin"

type FoodItem struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func ListFoods(c *gin.Context) {
	res := []FoodItem{
		{ID: 1, Name: "Pizza"},
		{ID: 2, Name: "Burger"},
		{ID: 3, Name: "Pasta"},
	}

	c.JSON(200, gin.H{"data": res})
}
