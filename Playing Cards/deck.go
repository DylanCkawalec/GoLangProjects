package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// creating a new type of deck
//which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven",
		"Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	//we always replace with an underscore because we dont need the index to show up after every iteration.
	for _, suit := range cardSuits {

		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	//return the cards after ranging through the values and suits
	return cards

}

//reference the Dekc type, and giving it a function whenever it's called named print.
func (d deck) print() {

	for i, card := range d {
		fmt.Println(i+1, card)
	}
}

//lets hand out some cards'
// we can can call this with type deck and a second with type int; and no other.
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

//create a helper function to turn decks to strings, for saving and turning strings to bytes.
//cards.toString makes sense.
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
	//from deck to [] string : type conversion ([]TypeWeWant("TypeWeHave"))

}

//now we have a way to convert our slices of []strings, to a single string.
// we have to prepare to WriteFile(filename string, data []yte, perm os.Filemode) error
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

//now that we saved a text file as a string, if we ever want to load it back to the computer
//we need to ReadFile(filename string)([]byte, error)
func newDeckFromFILE(filename string) deck {
	//the string object separated by commas is bs "bite slice"
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		//"Option#1" - log the err and return a call to newDeck()
		// :) Option#2 - Log the err and quit the program
		fmt.Println("Error: new deck failed to load properly", err)
		os.Exit(1)
	}
	//type conversion
	stringSlice := strings.Split(string(bs), ",")
	return deck(stringSlice)

}

func (d deck) shuffle() {
	//new source to source is like making a seed to new seed
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
