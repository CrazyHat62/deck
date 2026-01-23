// Package deck simply provides a 52 card deck shuffled with microsoft 32000 numbers (for freecell)
//
//	it also provides an :
//		additional 1 pass shuffle
//		ability to remove top or bottom card
//		ability to add top or bottom card
//		ability to split the deck in two
//
// decks are passed by slices so other things like a hand can be achieved
package deck

import (
	"errors"
	"fmt"
	"math"
	"math/rand/v2"
	//rl "github.com/gen2brain/raylib-go/raylib"
)

var Suits []string = []string{"hearts", "diamonds", "clubs", "spades"}
var Ranks []string = []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}
var Values = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}

type Card struct {
	Suit  string
	Rank  string
	index int
}

func NoCard() Card {
	return Card{Suit: "", Rank: "", index: -1}
}

// credit to https://rosettacode.org/wiki/Deal_cards_for_FreeCell#Go
const sSuits = "CDHS"
const sNums = "A23456789TJQK"

const rMax32 = math.MaxInt32

// credit to https://rosettacode.org/wiki/Deal_cards_for_FreeCell#Go
var seed = 1

// Used specifically for the Microsoft Deck numbers (first 32000)
// credit to https://rosettacode.org/wiki/Deal_cards_for_FreeCell#Go
func rnd() int {
	seed = (seed*214013 + 2531011) & rMax32
	return seed >> 16
}

// Create a deck and shuffle
// credit to https://rosettacode.org/wiki/Deal_cards_for_FreeCell#Go
func Deal(s int) []Card {
	seed = s
	t := make([]Card, 52)
	for i := 0; i < 52; i++ {
		c := 51 - i
		t[i].index = c
		t[i].Rank = fmt.Sprintf("%c", sNums[c/4])
		t[i].Suit = fmt.Sprintf("%c", sSuits[c%4])
	}
	for i := 0; i < 51; i++ {
		j := 51 - rnd()%(52-i)
		t[i], t[j] = t[j], t[i]
	}
	return t
}

// Shuffle(cards[:]) this will shuffle full or partial deck - different from above
func Shuffle(cards []Card) {
	for i := len(cards) - 1; i > 0; i-- {
		j := rand.IntN(len(cards))
		cards[i], cards[j] = cards[j], cards[i]
	}
}

// Color attached to card uses suit to determine if red or black ~ consider generalizing it
func (c Card) Color() string {
	if c.Suit == "hearts" || c.Suit == "diamonds" {
		return "red"
	} else {
		return "black"
	}
}

// PushLast will place the card at the bottom of the cards
// e.g. cards = PushLast(cards[:], c)
func PushLast(cards []Card, c Card) []Card {
	cards = append(cards, c)
	println("Pushed", c.Color(), c.Suit, c.Rank)
	return cards
}

// PushFirst will place the card at the Top of the cards
// e.g. cards = PushFirst(cards[:], c)
func PushFirst(cards []Card, c Card) []Card {
	cards = append([]Card{c}, cards...)
	return cards
}

// e.g. c, cards := PopFirst(cards[:])
func PopFirst(cards []Card) (Card, []Card, error) {
	if len(cards) > 0 {
		x := cards[0]
		cards = cards[1:]
		return x, cards, nil
	}
	return NoCard(), nil, errors.New("No cards in Deck")
}

// c, cards := PopLast(cards[:])
func PopLast(cards []Card) (Card, []Card, error) {
	if len(cards) > 0 {
		x := cards[len(cards)-1]
		cards = cards[:len(cards)-1]
		//println("returning", x.Color(), x.suite, x.CardValue())
		return x, cards, nil
	}
	return NoCard(), nil, errors.New("No cards in Deck")
}

func Split(cards []Card) ([]Card, []Card) {
	return cards[:int(len(cards)/2)], cards[int(len(cards)/2):]
}

func Show(cs []Card) {
	for i, c := range cs {
		fmt.Printf("%s%s", c.Rank, c.Suit)
		if (i+1)%8 == 0 || i+1 == len(cs) {
			fmt.Println()
		}
	}
}
