package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"slices"

	"backend/database"
)

func OpenAndUnmarshal(fileName string, v any) {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		log.Fatal(fmt.Sprintf("Could not open %s", fileName))
	}
	defer jsonFile.Close()
	data, err := io.ReadAll(jsonFile)
	err = json.Unmarshal(data, v)
	if err != nil {
		log.Fatal(fmt.Sprintf("Could not unmarshal %s: %v", fileName, err))
	}
}

func AssignIdUsers(users []*database.User) {
	for id, user := range users {
		user.Id = uint32(id)
	}
}

func AssignIdRecipes(recipes []*database.Recipe) {
	for id, recipe := range recipes {
		recipe.Id = uint32(id)
	}
}

func AssignIdIngredients(ingredients []*database.Ingredient) {
	for id, ingredient := range ingredients {
		ingredient.Id = uint32(id)
	}
}

func ConvertRecipesToShortRecipes(recipes []*database.RecipeWithProperties) []*database.RecipeShort {
	if len(recipes) == 0 {
		return nil
	}
	shortRecipes := make([]*database.RecipeShort, 0, len(recipes))
	for _, recipe := range recipes {
		shortRecipes = append(shortRecipes, &database.RecipeShort{
			Id:           recipe.Id,
			Name:         recipe.Name,
			Image:        recipe.Image,
			Rating:       recipe.Rating,
			Tags:         recipe.Tags,
			IconTags:     recipe.IconTags,
			CookingLevel: recipe.RecipeProperties.CookingLevel,
			CookingTime:  recipe.RecipeProperties.CookingTime,
		})
	}
	return shortRecipes
}

func FilterRecipesByTags(recipes []*database.RecipeWithProperties, user *database.User) []*database.RecipeWithProperties {
	retRecipes := make([]*database.RecipeWithProperties, len(recipes))
	copy(retRecipes, recipes)
	return slices.DeleteFunc(retRecipes, func(rwp *database.RecipeWithProperties) bool {
		ret := false
		ret = ret || (!rwp.Properties.AlcoholFree && user.HardRequirements.AlcoholFree)
		ret = ret || (!rwp.Properties.EcoFriendly && user.HardRequirements.EcoFriendly)
		ret = ret || (!rwp.Properties.EggFree && user.HardRequirements.EggFree)
		ret = ret || (!rwp.Properties.GlutenFree && user.HardRequirements.GlutenFree)
		ret = ret || (!rwp.Properties.LactoseFree && user.HardRequirements.LactoseFree)
		ret = ret || (!rwp.Properties.Mild && user.HardRequirements.Mild)
		ret = ret || (!rwp.Properties.MustardFree && user.HardRequirements.MustardFree)
		ret = ret || (!rwp.Properties.NutFree && user.HardRequirements.NutFree)
		ret = ret || (!rwp.Properties.PorkFree && user.HardRequirements.PorkFree)
		ret = ret || (!rwp.Properties.SoyFree && user.HardRequirements.SoyFree)
		ret = ret || (!rwp.Properties.Vegan && user.HardRequirements.Vegan)
		ret = ret || (!rwp.Properties.Vegetarian && user.HardRequirements.Vegetarian)
		ret = ret || (!rwp.Properties.WheatFree && user.HardRequirements.WheatFree)
		return ret
	})
}
