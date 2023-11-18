package main

import (
	"fmt"
	"log"
	"net/http"
	"slices"
	"strconv"

	"github.com/gin-gonic/gin"

	"backend/database"
	"backend/utils"
)

var (
	users       []*database.User
	ingredients []*database.Ingredient
	recipes     []*database.RecipeWithProperties
	tags        []string
)

func main() {
	utils.OpenAndUnmarshal("database/users.json", any(&users))
	utils.AssignIdUsers(users)
	utils.OpenAndUnmarshal("database/ingredients.json", any(&ingredients))
	utils.AssignIdIngredients(ingredients)
	var recipesTemp []*database.Recipe
	utils.OpenAndUnmarshal("database/recipes.json", any(&recipesTemp))
	utils.AssignIdRecipes(recipesTemp)
	recipes = utils.RecipesToRecipesWithProperties(recipesTemp, ingredients)
	utils.SetUserPreferences(users, tags)
	utils.AddIconTagsToRecipes(recipes)
	tags = utils.GetTags(recipes)

	router := gin.Default()
	router.Use(CORSMiddleware())
	router.GET("/recipes/:userId", getRecipes)
	router.GET("/tags", getTags)
	router.GET("/recipe/:id/:userId", getRecipe)
	router.GET("/image/:path", getImage)
	router.GET("/user/:id", getUser)
	router.GET("/user", getNewUserId)
	router.GET("/like/:recipeId/:userId", likeMeal)
	router.GET("/dislike/:recipeId/:userId", dislikeMeal)
	router.PUT("/user/:id/preferences", updateUserPreferences)
	router.PUT("/user/:id/requirements", updateUserRequirements)
	log.Fatal(router.Run(":8080"))
}

func getRecipes(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("userId"), 10, 32)
	if int(id) >= len(users) {
		ctx.JSON(http.StatusNotFound, nil)
		return
	}
	validRecipes := utils.FilterRecipesByTags(recipes, users[id])
	slices.SortFunc(validRecipes, func(i, j *database.RecipeWithProperties) int {
		sumPreferencesI := utils.SumPreferences(users[id].Preferences, i.Tags) / len(i.Tags)
		sumPreferencesJ := utils.SumPreferences(users[id].Preferences, j.Tags) / len(j.Tags)
		if sumPreferencesI == sumPreferencesJ {
			if i.Rating == j.Rating {
				return 0
			} else if i.Rating > j.Rating {
				return -1
			} else {
				return 1
			}
		} else if sumPreferencesI > sumPreferencesJ {
			return -1
		} else {
			return 1
		}
	})
	validRecipesShort := utils.ConvertRecipesToShortRecipes(validRecipes)
	ctx.JSON(http.StatusOK, validRecipesShort)
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
	userId, _ := strconv.ParseInt(ctx.Param("userId"), 10, 32)
	recipe := *recipes[id]
	if _, ok := users[userId].DislikedMeals[uint32(id)]; ok {
		recipe.Disliked = true
	}
	if _, ok := users[userId].LikedMeals[uint32(id)]; ok {
		recipe.Liked = true
	}
	ctx.JSON(http.StatusOK, recipe)
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

func likeMeal(ctx *gin.Context) {
	recipeId, _ := strconv.ParseInt(ctx.Param("recipeId"), 10, 32)
	userId, _ := strconv.ParseInt(ctx.Param("recipeId"), 10, 32)
	if _, ok := users[userId].LikedMeals[uint32(recipeId)]; !ok {
		users[userId].LikedMeals[uint32(recipeId)] = struct{}{}
		utils.AddPreferenceValue(users[userId], recipes[recipeId].Tags, 20)
	} else {
		delete(users[userId].LikedMeals, uint32(recipeId))
		utils.AddPreferenceValue(users[userId], recipes[recipeId].Tags, -20)
	}
}

func dislikeMeal(ctx *gin.Context) {
	recipeId, _ := strconv.ParseInt(ctx.Param("recipeId"), 10, 32)
	userId, _ := strconv.ParseInt(ctx.Param("recipeId"), 10, 32)
	if _, ok := users[userId].DislikedMeals[uint32(recipeId)]; !ok {
		users[userId].DislikedMeals[uint32(recipeId)] = struct{}{}
		utils.AddPreferenceValue(users[userId], recipes[recipeId].Tags, -20)
	} else {
		delete(users[userId].DislikedMeals, uint32(recipeId))
		utils.AddPreferenceValue(users[userId], recipes[recipeId].Tags, 20)
	}
}

func updateUserPreferences(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if int(id) >= len(users) {
		ctx.JSON(http.StatusNotFound, nil)
		return
	}
	var preferences map[string]map[string]bool
	err := ctx.BindJSON(&preferences)
	if err != nil {
		log.Print(err)
	}
	newPreferences := utils.ConvertPreferences(preferences["preferences"])
	utils.UpdatePreferences(newPreferences, users[id])
	ctx.JSON(http.StatusOK, nil)
}

func updateUserRequirements(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if int(id) >= len(users) {
		ctx.JSON(http.StatusNotFound, nil)
		return
	}
	var hardRequirements map[string]database.Properties
	err := ctx.ShouldBind(&hardRequirements)
	if err != nil {
		log.Print(err)
	}
	users[id].HardRequirements = hardRequirements["hardRequirements"]
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
