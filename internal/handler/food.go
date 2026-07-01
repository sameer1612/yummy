package handler

import "github.com/gin-gonic/gin"

type FoodItem struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Caption string  `json:"caption"`
	Rating  float32 `json:"rating"`
}

func ListFoods(c *gin.Context) {
	res := []FoodItem{
		{ID: 1, Name: "Pizza", Caption: "Delicious corn, onion and molten cheese pizza", Rating: 4.0},
		{ID: 2, Name: "Burger", Caption: "Crispy burger king potato patty burger", Rating: 4.5},
		{ID: 3, Name: "Pasta", Caption: "Creamy red and white mix sauce pasta", Rating: 4.8},
	}

	c.JSON(200, gin.H{"data": res})
}
