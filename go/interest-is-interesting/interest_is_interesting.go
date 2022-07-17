package interest

const (
	PenaltyPercentage float32 = 3.213
	lowInterest       float32 = .5
	mediumInterest    float32 = 1.621
	highInterest      float32 = 2.475
)

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) (Rate float32) {
	switch {
	case balance < 0:
		Rate = PenaltyPercentage
	case balance >= 0 && balance < 1000:
		Rate = lowInterest
	case balance >= 1000 && balance < 5000:
		Rate = mediumInterest
	case balance >= 5000:
		Rate = highInterest
	}
	return Rate
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	rate := InterestRate(balance) / 100
	return balance * float64(rate)
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return Interest(balance) + balance
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	var years int
	for i := 1; balance < targetBalance; i++ {
		balance = AnnualBalanceUpdate(balance)
		years = i
	}
	return years
}
