package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
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
        case "ten","jack","king","queen":
        	return 10
        default:
        	return 0
    }
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {

    playerPower := ParseCard(card1)+ParseCard(card2)
    dealerPower := ParseCard(dealerCard)
    switch {
        case playerPower == 22:
        	return "P"
        case playerPower == 21:
        	if dealerPower != 11 && dealerPower != 10{
                return "W"
            }
        	return "S"
        case playerPower >= 17 && playerPower <= 20:
        	return "S"
        case playerPower >= 12 && playerPower <= 16:
        	if dealerPower >= 7 {
                return "H"
            }
        	return "S"
        default:
        	return "H"
    }
}
