package utils

import (
	"fmt"
	"log"
	"slices"
	"strings"

	"backend/database"
	"backend/database/tags"
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
			return strings.Trim(strings.ToLower(ingredientName), " \n\t") == strings.Trim(strings.ToLower(ingredient.Name), " \n\t")
		}) {
			log.Fatal(fmt.Sprintf("Ingredient %s not found", ingredientName))
		}
	}
	return property
}

func AddIconTagsToRecipes(recipes []*database.RecipeWithProperties) {
	for _, recipe := range recipes {
		recipe.IconTags = getIconTags(recipe.Properties)
	}
}

func getIconTags(properties database.Properties) []string {
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

func ConvertPreferences(preferences map[string]bool) map[string]int {
	ret := make(map[string]int)

	for k, v := range preferences {
		if v {
			ret[k] = 2
		}
	}

	return ret
}

func UpdatePreferences(preferences map[string]int, user *database.User) {
	for k, v := range preferences {
		user.Preferences[k] *= v
	}
}

func SumPreferences(preferences map[string]int, tags []string) int {
	ret := 0
	for _, tag := range tags {
		ret += int(preferences[tag])
	}
	return ret
}
