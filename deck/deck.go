package deck

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
)

// Suit can be Clubs, Diamonds, Hearts or Spades
type Suit struct{ Name, Symbol string }

// Facevalue will be 2 to 10, Jack, Queen, King and Ace
type Facevalue struct {
	Name  string
	Value int
}

// Card is composed of a Suit and a Facevalue
type Card struct {
	Suit  Suit
	Value Facevalue
}

// Deck is a collection of cards.  There's no inherent limit on the size, but
// every time it's initialized, it starts with 52 cards.
type StandardDeck struct {
	Cards []Card
}

// GreaterThan tells if one card's Value is GreaterThan the other
func (card *Card) GreaterThan(b *Card) bool {
	return card.Value.Value > b.Value.Value
}

// LessThan tells if one card's Value is LessThan the other
func (card *Card) LessThan(b *Card) bool {
	return card.Value.Value < b.Value.Value
}

// Equal returns true or false based on the Facevalue
func (card *Card) Equal(b *Card) bool {
	return card.Value.Value == b.Value.Value
}

// Facecard returns true for J,Q,K,A, false for all others
func (card *Card) Facecard() (ans bool) {
	n := card.Value.Name
	return n == "Jack" || n == "Queen" || n == "King" || n == "Ace"
}

// ToStr returns a pretty string for the Cards
func (card *Card) ToStr() string {
	if card.Facecard() {
		return fmt.Sprintf(" %c%s", card.Value.Name[0], card.Suit.Symbol)
	}
	return fmt.Sprintf("%2d%s", card.Value.Value, card.Suit.Symbol)
}

// Initialize Initializes a deck of 52 Cards
func (deck *StandardDeck) Initialize() error {

	suits := []Suit{
		Suit{"Clubs", "♣"},
		Suit{"Diamonds", "♦"},
		Suit{"Hearts", "♥"},
		Suit{"Spades", "♠"},
	}
	facevalues := []Facevalue{
		Facevalue{"Two", 2},
		Facevalue{"Three", 3},
		Facevalue{"Four", 4},
		Facevalue{"Five", 5},
		Facevalue{"Six", 6},
		Facevalue{"Seven", 7},
		Facevalue{"Eight", 8},
		Facevalue{"Nine", 9},
		Facevalue{"Ten", 10},
		Facevalue{"Jack", 11},
		Facevalue{"Queen", 12},
		Facevalue{"King", 13},
		Facevalue{"Ace", 14},
	}

	// Empty the deck
	deck.Cards = nil

	// Fill the deck
	for _, suit := range suits {
		for _, facevalue := range facevalues {
			deck.Cards = append(deck.Cards, Card{Suit: suit, Value: facevalue})
		}
	}
	deck.Shuffle()
	return nil
}

// Shuffle randomizes a Deck using crypto/rand
func (deck *StandardDeck) Shuffle() (err error) {
	var old []Card
	old = deck.Cards
	var shuffled []Card
	// This should be a relatively fast shuffle, as we always pick a random number
	// within the remaining cards left to be shuffled.  This has the added bonus
	// of allowing a Standard deck to be a 'shoe' like in a casino where many Standard decks are
	// shuffled together and drawn from.

	// For N times (where N is the total number of cards in the deck)
	for i := len(old); i > 0; i-- {
		// Pick an index within the old cards
		nBig, e := rand.Int(rand.Reader, big.NewInt(int64(i)))
		if e != nil {
			panic(e)
		}
		j := nBig.Int64()
		// Append the randomly picked card to the 'shuffled' pile
		shuffled = append(shuffled, old[j])
		// remove the chosen card from the old pile and collapse
		// (length will be decremented)
		old = remove(old, j)
	}
	deck.Cards = shuffled
	return nil
}

// remove removes a card at index i from a slice of Cards and collapses the hole
// (length with be decremented)
func remove(slice []Card, i int64) []Card {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

// Draw will remove cards from the beginning of the deck (index 0)
// and return a slice of []Card type.
func (deck *StandardDeck) Draw(count int) (cards []Card, err error) {
	if count > len(deck.Cards) {
		return nil, errors.New("Not enough cards left in the deck")
	}

	hand := deck.Cards[0:count]
	deck.Cards = deck.Cards[count:]
	return hand, nil
}

// CardsLeft returns the number of cards left in the deck
func (deck *StandardDeck) CardsLeft() int {
	return len(deck.Cards)
}
