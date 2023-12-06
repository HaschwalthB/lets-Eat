package main

import (
	"github.com/gin-gonic/gin"
)

type Recipe struct {
	Name          string    `json:"name"`
	Tags          []string  `json:"tags"`
	Ingredients   []string  `json:"ingredients"`
	Instructions  []string  `json:"instructions"`
	PublishesdAtt time.Time`json:"publishedAt"`
}

var recipes []Recipe

func init() {
	recipes = make([]Recipe, 0)
}

func NewRecipe(c *gin.Context) {
	var recipe Recipe
	// marshal incoming request to Recipe 
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	recipe.ID = xid.New().String()
	recipe.PublishesdAtt = time.Now()
	recipes = append(recipes, recipe)
	c.JSON(http.StatusOK, recipe)
	PublishesdAtt time.Time `json:"publishedAt"`
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
}
