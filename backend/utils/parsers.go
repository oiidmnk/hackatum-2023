package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

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
