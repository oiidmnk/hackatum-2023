package database

type User struct {
	Id               uint32
	Name             string            `json:"name"`
	HardRequirements Properties        `json:"hard_requirements"`
	Preferences      map[string]uint32 `json:"preferences"`
}

type Properties struct {
	Vegan       bool `json:"vegan"`
	Vegetarian  bool `json:"vegetarian"`
	AlcoholFree bool `json:"alcohol_free"`
	MustardFree bool `json:"mustard_free"`
	LactoseFree bool `json:"lactose_free"`
	EggFree     bool `json:"egg_free"`
	PorkFree    bool `json:"pork_free"`
	WheatFree   bool `json:"wheat_free"`
	SoyFree     bool `json:"soy_free"`
	Mild        bool `json:"mild"`
	EcoFriendly bool `json:"eco_friendly"`
	GlutenFree  bool `json:"gluten_free"`
	NutFree     bool `json:"nut_free"`
}

type RecipeProperties struct {
	CookingTime  uint8    `json:"cooking_time"`
	Region       []string `json:"region"`
	CookingLevel uint8    `json:"cooking_level"`
}

type Ingredient struct {
	Id         uint32
	Name       string     `json:"name"`
	Properties Properties `json:"properties" json:"properties"`
}

type Recipe struct {
	Id               uint32
	Name             string           `json:"name"`
	Ingredients      []string         `json:"ingredients"`
	RecipeProperties RecipeProperties `json:"recipe_properties"`
	Tags             []string         `json:"tags"`
	Description      string           `json:"description"`
	Image            string           `json:"image"`
	Rating           uint8            `json:"rating"`
}

type RecipeWithProperties struct {
	Id               uint32
	Name             string
	Properties       Properties
	RecipeProperties RecipeProperties
	Tags             []string
	Description      string
	Image            string
	Rating           uint8
}

type RecipeShort struct {
	Id     uint32
	Name   string
	Image  string
	Rating uint32
	Tags   []string
}
