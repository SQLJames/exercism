package lasagnaMaster

const (
	PreperationTimePerLayerDefault = 2
	gramsOfNoodlesPerLayer = 50
	litersOfSaucePerLayer = .2
)

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string,PreperationTimePerLayer int) (PreparationTime int) {
	if PreperationTimePerLayer == 0 {
		PreperationTimePerLayer = PreperationTimePerLayerDefault
	}
	return len(layers) * PreperationTimePerLayer
}
// TODO: define the 'Quantities()' function
func Quantities(layers []string) (gramsOfNoodles int, litersofSauce float64)  {
	for _, v := range layers {
		if v == "noodles" {
			gramsOfNoodles += gramsOfNoodlesPerLayer
		}
		if v == "sauce" {
			litersofSauce += litersOfSaucePerLayer
		}
	}
	return gramsOfNoodles,litersofSauce
}
// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string)  {
	myList[len(myList)-1] = friendsList[len(friendsList)-1] 
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amounts []float64, portions int) (ScaledAmounts []float64) {
	for _, v := range amounts {
		ScaledAmounts = append(ScaledAmounts, v/2 * float64(portions))
	}
	return ScaledAmounts
}
