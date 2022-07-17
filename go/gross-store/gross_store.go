package gross

type units map[string]int

var (
	measurements = units{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
)

// Units stores the Gross Store unit measurements.
func Units() (Measurements map[string]int) {
	return measurements
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int, 5)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	if !unitExists(units, unit) {
		return false
	}
	if !itemInBill(bill, item) {
		bill[item] = units[unit]
	} else {
		bill[item] = units[unit] + bill[item]
	}
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	if !unitExists(units, unit) || !itemInBill(bill, item) {
		return false
	}
	quantityOnBill := bill[item]
	quantityToRemove := units[unit]
	quantityLeftOver := quantityOnBill - quantityToRemove
	if quantityLeftOver < 0 {
		return false
	}
	if quantityLeftOver == 0 {
		delete(bill, item)
	}
	if quantityLeftOver > 0 {
		bill[item] = quantityLeftOver
	}
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	if itemInBill(bill, item) {
		return bill[item], true
	}
	return 0, false
}

func unitExists(units map[string]int, unit string) bool {
	_, unitExists := units[unit]
	return unitExists
}

func itemInBill(bill map[string]int, item string) bool {
	_, itemInBill := bill[item]
	return itemInBill
}
