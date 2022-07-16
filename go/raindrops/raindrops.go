package raindrops

import "strconv"

const (
	divisibleByThree string = "Pling"
	divisibleByFive string = "Plang"
	divisibleBySeven string = "Plong"
)

func Convert(number int) string {
	var value string
	if number % 3 == 0 {
		value += divisibleByThree
	}
	if number % 5 == 0 {
		value += divisibleByFive
	}
	if number % 7 == 0 {
		value += divisibleBySeven
	}
	if value == "" {
		value = strconv.Itoa(number)
	}
	return value
}
