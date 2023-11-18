package utils

import (
	"backend/database"
	"backend/database/tags"
	"fmt"
	"log"
	"slices"
	"strings"
)

func RecipesToRecipesWithProperties(recipes []*database.Recipe, ingredients []*database.Ingredient) []*database.RecipeWithProperties {
	recipesWithProperties := make([]*database.RecipeWithProperties, 0, len(recipes))
	for _, recipe := range recipes {
		recipesWithProperties = append(recipesWithProperties, &database.RecipeWithProperties{
			Id:                  recipe.Id,
			Name:                recipe.Name,
			Properties:          ingredientsToProperties(recipe.Ingredients, ingredients),
			RecipeProperties:    recipe.RecipeProperties,
			Tags:                recipe.Tags,
			Description:         recipe.Description,
			Image:               recipe.Image,
			Rating:              recipe.Rating,
			CookingInstructions: recipe.CookingInstructions,
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
	for _, ingredient := range allIngredients {
		if slices.ContainsFunc(ingredientNames, func(ingredientName string) bool {
			return strings.Trim(strings.ToLower(ingredientName), " \n\t") == strings.Trim(strings.ToLower(ingredient.Name), " \n\t")
		}) {
			property.Vegan = ingredient.Properties.Vegan && property.Vegan
			property.Vegetarian = ingredient.Properties.Vegetarian && property.Vegetarian
			property.AlcoholFree = ingredient.Properties.AlcoholFree && property.AlcoholFree
			property.MustardFree = ingredient.Properties.MustardFree && property.MustardFree
			property.LactoseFree = ingredient.Properties.LactoseFree && property.LactoseFree
			property.EggFree = ingredient.Properties.EggFree && property.EggFree
			property.PorkFree = ingredient.Properties.PorkFree && property.PorkFree
			property.WheatFree = ingredient.Properties.WheatFree && property.WheatFree
			property.SoyFree = ingredient.Properties.SoyFree && property.SoyFree
			property.Mild = ingredient.Properties.Mild && property.Mild
			property.EcoFriendly = ingredient.Properties.EcoFriendly && property.EcoFriendly
			property.GlutenFree = ingredient.Properties.GlutenFree && property.GlutenFree
			property.NutFree = ingredient.Properties.NutFree && property.NutFree
		}
	}
	for _, ingredientName := range ingredientNames {
		if !slices.ContainsFunc(allIngredients, func(ingredient *database.Ingredient) bool {
			return ingredient.Name == ingredientName
		}) {
			log.Fatal(fmt.Sprintf("Ingredient %s not found", ingredientName))
		}
	}
	return property
}

func AddTagsToRecipes(recipes []*database.RecipeWithProperties) {
	for _, recipe := range recipes {
		recipe.Tags = append(recipe.Tags, getTags(recipe.Properties)...)
	}
}

func getTags(properties database.Properties) []string {
	ret := make([]string, 0)

	if properties.Vegan {
		ret = append(ret, tags.VEGAN)
	}
	if properties.Vegetarian {
		ret = append(ret, tags.VEGETARIAN)
	}
	if properties.GlutenFree {
		ret = append(ret, tags.GLUTENFREE)
	}
	if !properties.Mild {
		ret = append(ret, tags.SPICY)
	}

	return ret
}

func GetTags(recipes []*database.RecipeWithProperties) []string {
	ret := make(map[string]bool, 0)

	for _, recipe := range recipes {
		for _, tag := range recipe.Tags {
			ret[tag] = true
		}
	}

	keys := make([]string, len(ret))

	i := 0
	for k := range ret {
		keys[i] = k
		i++
	}

	return keys
}
