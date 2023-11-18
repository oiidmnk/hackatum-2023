package main

import (
	"fmt"

	"backend/database"
	"backend/utils"
)

func main() {
	var users []*database.User
	utils.OpenAndUnmarshal("database/users.json", any(&users))
	utils.AssignIdUsers(users)
	var ingredients []*database.Ingredient
	utils.OpenAndUnmarshal("database/ingredients.json", any(&ingredients))
	utils.AssignIdIngredients(ingredients)
	var recipes []*database.Recipe
	utils.OpenAndUnmarshal("database/recipes.json", any(&recipes))
	utils.AssignIdRecipes(recipes)
	recipesShort := utils.ConvertRecipesToShortRecipes(recipes)
	fmt.Println(recipesShort[0].Name)

	return
}
