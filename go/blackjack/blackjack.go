package blackjack

type deck struct {
	cards map[string]int
}

func newDeck() deck {
	cards := map[string]int{
		"ace" :11,
		"two" :2,
		"three" :3,
		"four" :4,
		"five" :5,
		"six" :6,
		"seven" :7,
		"eight" :8,
		"nine" :9,
		"ten" :10,
		"jack" :10,
		"queen" :10,
		"king" :10,
	}
	return deck{cards: cards}
}

var (
	d = newDeck()
)

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	value, ok := d.cards[card]
	if !ok {
		return 0
	}
	return value
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	handTotal := ParseCard(card1) + ParseCard(card2)
	DealerShowing := ParseCard(dealerCard)
	//If you have a pair of aces you must always split them.
	if card1 == "ace" && card2 == "ace" {
		return "P"
	}
	//If your cards sum up to 11 or lower you should always hit.
	if handTotal <= 11 {
		return "H"
	}
	//unless the dealer has a 7 or higher, in which case you should always hit.
	if (handTotal >= 12 && handTotal <= 16) && DealerShowing >= 7 {
		return "H"
	}
	//If you have a Blackjack (two cards that sum up to a value of 21),
	//and the dealer does not have an ace, a figure or a ten then you automatically win.
	if handTotal == 21 && !(DealerShowing == 10 || DealerShowing == 11) {
		return "W"
	}
	//If the dealer does have any of those cards then you'll have to stand and wait for the reveal of the other card.
	return "S"
}
