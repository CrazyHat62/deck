package deck

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var ClubsSprites rl.Texture2D
var SpadesSprites rl.Texture2D
var DiamondsSprites rl.Texture2D
var HeartsSprites rl.Texture2D

const SpriteSizeW = 125 //500
const SpriteSizeH = 181 //726

func LoadStandardDeckSprites() {
	ClubsSprites = rl.LoadTexture("images/clubsS.png")
	SpadesSprites = rl.LoadTexture("images/spadesS.png")
	DiamondsSprites = rl.LoadTexture("images/diamondsS.png")
	HeartsSprites = rl.LoadTexture("images/heartsS.png")
}
func UnloadStandardDeckSprites() {
	defer rl.UnloadTexture(ClubsSprites)
	defer rl.UnloadTexture(SpadesSprites)
	defer rl.UnloadTexture(DiamondsSprites)
	defer rl.UnloadTexture(HeartsSprites)
}

// returns the source rectangle for the card sprite sheet based on the card's rank and suit
func GetCardSource(card Card, frameRec rl.Rectangle) rl.Rectangle {
	ix := int32(0)
	iy := int32(0)
	switch { //could use dictionary to map
	case card.Rank == string('2'):
		ix = 0
		iy = 0
	case card.Rank == string('3'):
		ix = 1
		iy = 0
	case card.Rank == string('4'):
		ix = 2
		iy = 0
	case card.Rank == string('5'):
		ix = 3
		iy = 0
	case card.Rank == string('6'):
		ix = 4
		iy = 0
	case card.Rank == string('7'):
		ix = 5
		iy = 0
	case card.Rank == string('8'):
		ix = 6
		iy = 0
	case card.Rank == string('9'):
		ix = 7
		iy = 0
	case card.Rank == string('T'):
		ix = 0
		iy = 1
	case card.Rank == string('A'):
		ix = 1
		iy = 1
	case card.Rank == string('J'):
		ix = 5
		iy = 1
	case card.Rank == string('K'):
		ix = 6
		iy = 1
	case card.Rank == string('Q'):
		ix = 7
		iy = 1
	}

	source := frameRec
	source.X = float32(SpriteSizeW*ix + 1)
	source.Y = float32(SpriteSizeH * iy)

	return source
}
func GetSuitSprite(card Card, TxSprites rl.Texture2D) rl.Texture2D {
	switch {
	case card.Suit == string('C'):
		TxSprites = ClubsSprites
	case card.Suit == string('S'):
		TxSprites = SpadesSprites
	case card.Suit == string('D'):
		TxSprites = DiamondsSprites
	case card.Suit == string('H'):
		TxSprites = HeartsSprites
	}
	return TxSprites
}
