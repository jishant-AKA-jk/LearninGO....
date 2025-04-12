package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgTime int) int {
	if avgTime == 0 {
		avgTime = 2
	}
	return len(layers) * avgTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	totalNoodles := 0
	totalSauce := 0.0
	for i := 0; i < len(layers); i++ {
		if layers[i] == "noodles" {
			totalNoodles += 50
		}
		if layers[i] == "sauce" {
			totalSauce += 0.2
		}
	}
	return totalNoodles, totalSauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, scale int) []float64 {
	var newRecipe []float64
	for i := 0; i < len(quantities); i++ {
		newRecipe = append(newRecipe, (quantities[i] * (float64(scale) / 2)))
	}
	return newRecipe
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
