package cars

const (
	costOfTenCars = 95000
	costOfoneCar = 10000
)

func ratePercentage(Rate float64) (percentage float64) {
	return Rate / 100
}

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) (CarsPerHour float64) {
	return float64(productionRate) * ratePercentage(successRate)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate) / 60)
}

// CalculateCost works out the cost of producing the given number of cars
func CalculateCost(carsCount int) uint {
	groups := carsCount / 10
	remainder := carsCount % 10
	return uint((groups * costOfTenCars) + (remainder * costOfoneCar))
}
