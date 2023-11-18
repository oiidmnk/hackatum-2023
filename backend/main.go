package main

import (
	"backend/database"
	"backend/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

var (
	users        []*database.User
	ingredients  []*database.Ingredient
	recipes      []*database.RecipeWithProperties
	recipesShort []*database.RecipeShort
	tags         []string
)

func main() {
	utils.OpenAndUnmarshal("database/users.json", any(&users))
	utils.AssignIdUsers(users)
	utils.SetUserPreferences(users, tags)
	utils.OpenAndUnmarshal("database/ingredients.json", any(&ingredients))
	utils.AssignIdIngredients(ingredients)
	var recipesTemp []*database.Recipe
	utils.OpenAndUnmarshal("database/recipes.json", any(&recipesTemp))
	utils.AssignIdRecipes(recipesTemp)
	recipes = utils.RecipesToRecipesWithProperties(recipesTemp, ingredients)
	utils.AddTagsToRecipes(recipes)
	recipesShort = utils.ConvertRecipesToShortRecipes(recipes)
	tags = utils.GetTags(recipes)

	router := gin.Default()
	router.Use(CORSMiddleware())
	router.GET("/recipes", getRecipes)
	router.GET("/tags", getTags)
	router.GET("/recipe/:id", getRecipe)
	router.GET("/image/:path", getImage)
	router.GET("/user/:id", getUser)
	router.GET("/user", getNewUserId)
	router.POST("/user/:id/preferences", updateUserPreferences)
	router.POST("/user/:id/requirements", updateUserRequirements)
	log.Fatal(router.Run(":8080"))
}

func getRecipes(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, recipesShort)
}

func getImage(ctx *gin.Context) {
	name := ctx.Param("path")
	if name[len(name)-4:] == ".png" {
		ctx.Header("Content-Type", "image/png")
	} else if name[len(name)-4:] == ".jpg" || name[len(name)-5:] == ".jpeg" {
		ctx.Header("Content-Type", "image/jpeg")
	}
	ctx.File(fmt.Sprintf("database/images/%s", name))
}

func getTags(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, tags)
}

func getRecipe(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 32)
	ctx.JSON(http.StatusOK, recipes[id])
}

func getUser(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if int(id) >= len(users) {
		ctx.JSON(http.StatusNotFound, nil)
		return
	}
	ctx.JSON(http.StatusOK, int(id))
}

func getNewUserId(ctx *gin.Context) {
	id := uint32(len(users))
	users = append(users, &database.User{
		Id:               id,
		HardRequirements: database.Properties{},
		Preferences:      utils.CreatePreferences(nil, tags),
	})
	ctx.JSON(http.StatusOK, id)
}

func updateUserPreferences(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if int(id) >= len(users) {
		ctx.JSON(http.StatusNotFound, nil)
		return
	}
	var preferences map[string]uint32
	err := ctx.BindJSON(&preferences)
	if err != nil {
		log.Print(err)
	}
	users[id].Preferences = utils.CreatePreferences(preferences, tags)
	ctx.JSON(http.StatusOK, nil)
}

func updateUserRequirements(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if int(id) >= len(users) {
		ctx.JSON(http.StatusNotFound, nil)
		return
	}
	var hardRequirements database.Properties
	err := ctx.BindJSON(&hardRequirements)
	if err != nil {
		log.Print(err)
	}
	users[id].HardRequirements = hardRequirements
	ctx.JSON(http.StatusOK, nil)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, DELETE")
		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}
		ctx.Next()
	}
}
