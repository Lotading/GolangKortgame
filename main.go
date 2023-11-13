package main

import (
	"fmt"
)

type Deck struct {
	Kort []string
}

type Gamers struct {
	Antall int
}

func main() {
	deck := &Deck{}
	gamers := &Gamers{}

	hand(deck)
	spiller(gamers)
	KortGame(deck,gamers)
}

func hand(Deck *Deck) {
	KortDeck := []string{"1","2","3","4","5","6","7","8","9","10","J","Q","K"}
	Suits := []string{"Dia","Hje","Spa","Klø"}

	for i := 0; i < count; i++ {

		for i := 0; i < count; i++ {
			
		}

	}

	// Deck.Kort = Kort
	// fmt.Println(Kort)
}

func spiller(Gamers *Gamers) {

	var Antall int
	fmt.Print("Hvor mange spillere er det? ")
	fmt.Scan(&Antall)

	switch {
		case Antall <= 1:
			fmt.Println("???")
		case Antall > 8:
			fmt.Println("For mange spillere")
		default:
			Gamers.Antall = Antall
			return
	}
}

func KortGame(deck *Deck,Gamers *Gamers){
	
}
