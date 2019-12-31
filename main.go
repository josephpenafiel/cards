package main
// import "fmt"
import "github.com/josephpenafiel/cards/deck"

func main() {
	cards := deck.NewDeck()
	cards.SaveToFile("deck2.txt")
	cards.Shuffle()
}