// Package weather package gives you functions to get information about the weather in an area.
package weather

var (
	// CurrentCondition documents the weather outside currently.
	CurrentCondition string
	// CurrentLocation docouments the current location of the user.
	CurrentLocation string
)

// Forecast returns the weather forcast given a city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
