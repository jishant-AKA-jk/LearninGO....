package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
// ace	    11	eight	8
// two	    2	 nine	9
// three	3	ten	10
// four	4	jack	10
// five	 5	queen	10
// six	    6	king	10
// seven	7	other	0
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten":
		return 10
	case "jack":
		return 10
	case "queen":
		return 10
	case "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	var sum = ParseCard(card1) + ParseCard(card2)
	if sum == 22 {
		return "P"
	} else if sum == 21 {
		if ParseCard(dealerCard) < 10 {
			return "W"
		}
		return "S"
	} else if sum <= 20 && sum >= 17 {
		return "S"
	} else if sum <= 16 && sum >= 12 {
		if ParseCard(dealerCard) >= 7 {
			return "H"
		}
		return "S"
	} else {
		return "H"
	}

}
