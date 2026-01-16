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
    case "ten", "jack", "queen", "king": // You can comma-separate multiple cases!
        return 10
    default:
        return 0 // This handles "other" cards
    }
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    
    playerVal := ParseCard(card1) + ParseCard(card2)
    dealerVal := ParseCard(dealerCard)

    switch {
    case playerVal == 22:
        return "P"
    case playerVal == 21:
        if dealerVal >= 10 {
            return "S"
        } else {
            return "W"
        }
    case playerVal >= 17 && playerVal <= 20:
        return "S"
    case playerVal >= 12 && playerVal <= 16:
        if dealerVal >= 7 {
            return "H"
        } else {
            return "S"
        }
    case playerVal <= 11:
        return "H"
    }
    
	panic("Please implement the FirstTurn function")
}
