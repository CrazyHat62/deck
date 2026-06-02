package deck

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestDeal(t *testing.T) {

	got := Deal(2570)
	want := []string{"2S2", "KD13", "JD11", "8H8", "8S8", "QD12", "4C4", "3S3",
		"8D8", "AH1", "3C3", "5S5", "2D2", "7C7", "7D7", "QC12",
		"QS12", "6C6", "6D6", "5H5", "KH13", "KS13", "7S7", "9H9",
		"4D4", "2H2", "3D3", "9C9", "4S4", "QH12", "7H7", "6S6",
		"AC1", "3H3", "KC13", "TD10", "2C2", "4H4", "JC11", "TH10",
		"9S9", "5D5", "JH11", "TC10", "6H6", "9D9", "JS11", "AS1",
		"5C5", "8C8", "TS10", "AD1"}

	for i, c := range got {

		g := c.Rank + c.Suit + fmt.Sprint(c.Value)
		if g != want[i] {
			t.Errorf("got %q want %q", g, want[i])
		}

	}
	print(len(got))

}

func TestShuffle(t *testing.T) {

	gotSeed := 1 + rand.Intn(32000)
	wantSeed := 1 + rand.Intn(32000)

	got := Deal(gotSeed)
	want := Deal(wantSeed)
	//length must be the same
	if len(got) != len(want) {
		t.Errorf("got %d want %d", len(got), len(want))
		return
	}
OuterLoop:
	for _, w := range want {
		for _, g := range got {
			if w.Suit == g.Suit && w.Rank == g.Rank {
				continue OuterLoop
			}
		}
		t.Errorf("not found: wanted %q %q", w.Rank, w.Suit)
	}

	count := 0
	atMost := 5
	for i, c := range want {
		cg := got[i]
		if c.Suit == cg.Suit && c.index == cg.index {
			count++
		}
	}
	if count > atMost {
		t.Errorf("got %d want > %d", count, atMost)
	}

}

func TestColor(t *testing.T) {
	var c Card = Card{Suit: "hearts", index: 0}
	if c.Color() != "red" {
		t.Errorf("got %q want %q", c.Color(), "red")
	}
	c = Card{Suit: "spades", index: 13}
	if c.Color() != "black" {
		t.Errorf("got %q want %q", c.Color(), "black")
	}

}
