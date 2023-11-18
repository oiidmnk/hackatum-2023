package properties

const (
	NUTFREE     = "Nut free"
	MUSTARDFREE = "Mustard free"
	GLUTENFREE  = "Gluten free"
	LACTOSEFREE = "Lactose free"
	MILKFREE    = "Milk free"
	EGGFREE     = "Egg free"
	WHEATFREE   = "Wheat free"
	SOYFREE     = "Soy free"
	PORKFREE    = "Pork free"
	ALCOHOLFREE = "Alcohol free"
	VEGETARIAN  = "Vegetarian"
	VEGAN       = "Vegan"
	MILD        = "Mild"
	ECOFRIENDLY = "Eco friendly"
	NONE
)

func GetHardRequirements() []string {
	return []string{NUTFREE, MUSTARDFREE, GLUTENFREE, LACTOSEFREE, MILKFREE, EGGFREE, WHEATFREE, SOYFREE, PORKFREE, ALCOHOLFREE, VEGETARIAN, VEGAN, MILD, ECOFRIENDLY, NONE}
}
