package purchase

import "sort"

var (
	requiresLicense = [2]string{"car", "truck"}
)

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	for _, v := range requiresLicense {
		if v == kind {
			return true
		}
	}
	return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) (selection string) {
	options := []string{option1, option2}
	sort.Strings(options)
	return options[0] + " is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age < 3 {
		return originalPrice * .8
	} else if age > 3 && age < 10 {
		return originalPrice * .7
	}
	return originalPrice * .5
}
