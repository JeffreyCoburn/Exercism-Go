package lasagna

// PreparationTime accepts a slice of layers and the average preparation time per layer in minutes and returns an estimate for the total preparation time.
// If the average preparation time passed in is 0, then a default value of 2 is used.
func PreparationTime(layers []string, timePerLayer int) int {
	if timePerLayer == 0 {
		timePerLayer = 2
	}
	return len(layers) * timePerLayer
}

// Quantities takes a slice of layers and returns the amount of noodles needed in grams and the amount of sauce needed in liters
func Quantities(layers []string) (noodles int, sauce float64) {
	noodlesPerLayer := 50
	saucePerLayer := 0.2
	noodles = 0
	sauce = 0.0
	for _, layer := range layers {
		if layer == "noodles" {
			noodles += noodlesPerLayer
		}
		if layer == "sauce" {
			sauce += saucePerLayer
		}
	}
	return
}

// AddSecretIngredient takes a list of ingredients of your friend's recipe and the list of ingredients of your recipe,
// and adds the last recipe of their recipe to yours. The secret ingredient is assumed to be the last ingredient of their list
func AddSecretIngredient(friendsList []string, myList []string) []string {
	list := append(myList, friendsList[len(friendsList)-1])
	return list
}

// ScaleRecipe takes a slice of the amount of each ingredient needed to make 2 portions, the number of portions desired,
// and returns the amounts of each ingredient needed for the desired number of portions
func ScaleRecipe(amounts []float64, portions int) []float64 {
	scale := float64(portions) / 2.0
	scaledAmounts := make([]float64, len(amounts))
	for ingredient, amount := range amounts {
		scaledAmounts[ingredient] = amount * scale
	}
	return scaledAmounts
}
