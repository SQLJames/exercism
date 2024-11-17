package techpalace

import (
	"fmt"
	"strings"
)

const (
	lineTermination string = "\n"
)

func createStars(numStarsPerLine int) (stars string) {
	return strings.Repeat("*", numStarsPerLine)
}

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return fmt.Sprintf("Welcome to the Tech Palace, %s", strings.ToUpper(customer))
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) (WelcomeBorder string) {
	stars := createStars(numStarsPerLine)
	return stars + lineTermination + welcomeMsg + lineTermination + stars
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) (marketingMsg string) {
	marketingMsg = strings.ReplaceAll(oldMsg, "*", "")
	marketingMsg = strings.TrimSpace(marketingMsg)
	return marketingMsg
}
