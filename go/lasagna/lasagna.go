package lasagna

const (
	OvenTime                    int = 40
	PreparationTimeForEachLayer int = 2
)

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) (remaining int) {
	return OvenTime - actualMinutesInOven
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) (PrepTime int) {
	return PreparationTimeForEachLayer * numberOfLayers
}

// ElapsedTime calculates the total time needed to create and bake a lasagna.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) (TimeSpent int) {
	return PreparationTime(numberOfLayers) + actualMinutesInOven
}
