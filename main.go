package main

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

type Recipe struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	Tags          []string  `json:"tags"`
	Ingredients   []string  `json:"ingredients"`
	Instructions  []string  `json:"instructions"`
	PublishesdAtt time.Time `json:"publishedAt"`
}

var recipes []Recipe

func init() {
	recipes = make([]Recipe, 0)
  file, _ := os.ReadFile("recipes.json")
  _ = json.Unmarshal([]byte(file), &recipes)
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
}

func ListRecipes(c *gin.Context) {
  c.JSON(http.StatusOK, recipes)
}

func main() {
	r := gin.Default()
  r.POST("/recipes", NewRecipe)
  r.GET("/recipes", ListRecipes)
  r.Run()
}
