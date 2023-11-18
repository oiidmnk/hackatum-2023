package utils

import (
	"slices"

	"backend/database"
)

func RecipesToRecipesWithProperties(recipes []*database.Recipe, ingredients []*database.Ingredient) []*database.RecipeWithProperties {
	recipesWithProperties := make([]*database.RecipeWithProperties, 0, len(recipes))
	for _, recipe := range recipes {
		recipesWithProperties = append(recipesWithProperties, database.RecipeWithProperties{
			Id:               recipe.Id,
			Name:             recipe.Name,
			Properties:       ingredientsToProperties(recipe.Ingredients, ingredients),
			RecipeProperties: database.RecipeProperties{},
			Tags:             nil,
			Description:      "",
			Image:            "",
			Rating:           0,
		})
	}
	return recipesWithProperties
}

func ingredientsToProperties(ingredientNames []string, allIngredients []*database.Ingredient) database.Properties {
	property := database.Properties{
		Vegan:       true,
		Vegetarian:  true,
		AlcoholFree: true,
		MustardFree: true,
		LactoseFree: true,
		EggFree:     true,
		PorkFree:    true,
		WheatFree:   true,
		SoyFree:     true,
		Mild:        true,
		EcoFriendly: true,
		GlutenFree:  true,
		NutFree:     true,
	}
	ingredients := make([]*database.Properties, 0, len(allIngredients))
	for _, ingredient := range allIngredients {
		if slices.Contains(ingredientNames, ingredient.Name) {
			property.Vegan = ingredient.Properties.Vegan && property.Vegan
			property.Vegetarian = ingredient.Properties.Vegetarian && property.Vegetarian
			property.AlcoholFree = ingredient.Properties.AlcoholFree && property.AlcoholFree
			property.MustardFree = ingredient.Properties.MustardFree && property.MustardFree
			property.LactoseFree = ingredient.Properties.LactoseFree && property.LactoseFree
		}
	}
}
