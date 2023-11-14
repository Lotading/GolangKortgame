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

type spillerKort struct {
	spKort map[string]string
}

func main() {
	deck := &Deck{}
	gamers := &Gamers{}
	spillerkort := &spillerKort{}

	hand(deck)
	spiller(gamers)
	KortGame(deck,gamers,spillerkort)
}

func hand(Deck *Deck) {

	var kortStokk []string

	Kort := []string{"1","2","3","4","5","6","7","8","9","10","J","Q","K"}
	Suits := []string{"Dia","Hje","Spa","Klø"}

	for i := 0; i < len(Suits); i++ {
		for j := 0; j < len(Kort); j++ {
			card := Kort[j] + Suits[i]
			kortStokk = append(kortStokk, card)
		}
	}
	Deck.Kort = kortStokk
}

func spiller(Gamers *Gamers) {

	var Antall int
	fmt.Print("Hvor mange spillere er det? ")
	fmt.Scan(&Antall)

	switch {
		case Antall <= 1:
			fmt.Println("Vennligst få flere til å være med")
		case Antall > 8:
			fmt.Println("For mange spillere")
		default:
			Gamers.Antall = Antall
			return
	}
}

func KortGame(deck *Deck,Gamers *Gamers, spillerKort *spillerKort) {
	
}
