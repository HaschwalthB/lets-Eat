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

// all of the recipes
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

func UpdateRecipes(c *gin.Context) {
	id := c.Param("id")
	var recipe Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// intialize last index:
	// loop through recipes, and increase i based on len recipes
	index := -1
	for i := 0; i < len(recipes); i++ {
		if recipes[i].ID == id {
			index = i
		}
	}

	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "recipes not found",
		})
		return
	}

	recipes[index] = recipe
	c.JSON(http.StatusOK, recipe)
}

func DeleteRecipes(c *gin.Context) {
	id := c.Param("id")
	// intialize last index:
	index := -1
	for i := 0; i < len(recipes); i++ {
		if recipes[i].ID == id {
			index = i
		}
	}
	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "recipes not found",
		})
		return
	}
	// remove recipes from slice
  recipes = append(recipes[:index], recipes[index+1:]...)
  c.JSON(http.StatusOK, gin.H{
    "message" : "recipe deleted",
  })
}

func main() {
	r := gin.Default()
	r.POST("/recipes", NewRecipe)
	r.GET("/recipes", ListRecipes)
	r.PUT("/recipes/:id", UpdateRecipes)
	r.DELETE("/recipes/:id", DeleteRecipes)
	r.Run()
}
