package main

import (
	"fmt"
	"math/rand"
)

type Game struct {
	Deck []Card
}

type Card struct {
	Suit string
	Number string
}


func GenerateDeck() []Card {
	var deck []Card
	suits := []string{"Heart", "Diamond", "Club", "Spade"}
	numbers := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	for i:=0; i< len(suits); i++ {
		for j:=0; j< len(numbers); j++{
			card := Card{
				Suit: suits[i],
				Number: numbers[j],
			}
				fmt.Println(card)

				deck = append(deck, card)	
		}
	}
	// Generate a standard deck of cards (52 cards, 4 suits)
	return deck
}

func Shuffle(d []Card) []Card { 
	for j := 1; j < len(d); j++ {
		r := rand.Intn(j + 1)
		if j != r {
			d[r], d[j] = d[j], d[r]
		}
	}
	return d
}

func Deal(d []Card, n int) []Card {
	var deck []Card
	for i := 0; i < n; i++ {
		deck = append(deck, d[i])	
	}
	return deck
}

func main() {
	var numberOfCards = 10 // number of cards for each hand
	Deck:= GenerateDeck()
	deal:= Deal(Shuffle(Deck), numberOfCards * 2)

	playerOne := deal[0:numberOfCards]
	playerTwo := deal[numberOfCards:len(deal)]
	fmt.Println(playerOne)
	fmt.Println(playerTwo)
	// Deal two hands, print out both hands here
}