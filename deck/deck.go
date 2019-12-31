package deck

import (
	"time"
	"math/rand"
	"os"
	"fmt"
	"strings"
	"io/ioutil"
)

// Deck type
type Deck []string

//NewDeck creates new deck
func NewDeck() Deck {
	cards := Deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards

}


//Deal Gives cards to player
func Deal(handSize int, d *Deck) (Deck, Deck) { // returns a tuple

	hand := (*d)[:handSize]
	remainder := (*d)[handSize:]

	return hand, remainder
}

//Print Prints the created deck
func (d Deck) Print() {
	for card := range d {
		fmt.Println(card)
	}
}

func (d Deck) toString() string {
	deck := []string(d)
	return strings.Join(deck, "\n")
}

// SaveToFile Saves a deck to a file
func (d Deck) SaveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}


// NewDeckFromFile creates a deck from a .txt file
func NewDeckFromFile(filename string) Deck {
	byteSlice, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(byteSlice), "\n")
	return Deck(s)
}

// Shuffle Shuffles the current deck
func (d Deck) Shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}