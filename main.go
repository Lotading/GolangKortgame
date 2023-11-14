package main

import (
	"fmt"
	"math/rand"
)

type Deck struct {
	Kort []string
}

type Gamers struct {
	Antall  int
	navn    string
	navnArr []string
}

type spillerKort struct {
	spKort map[string][]string
}

func main() {
	deck := &Deck{}
	gamers := &Gamers{}
	spillerkort := &spillerKort{}

	hand(deck)
	spiller(gamers)
	shuffleNgive(deck, gamers, spillerkort)
}

func hand(Deck *Deck) {

	var kortStokk []string

	Kort := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	Suits := []string{"Dia", "Hje", "Spa", "Klø"}

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
		navn(Gamers)
	}

}

func navn(Gamers *Gamers) {
	var navn string
	var navnArr []string

	for i := 0; i < Gamers.Antall; i++ {
		fmt.Print("Hva heter de? ")
		fmt.Scan(&navn)
		navnArr = append(navnArr, navn)
	}

	Gamers.navnArr = navnArr
	Gamers.navn = navn
	fmt.Println(navnArr)
}

func shuffleNgive(deck *Deck, Gamers *Gamers, spillerKort *spillerKort) {

	spDeck := make(map[string][]string)
	
	spillere := Gamers.navnArr
	hand := deck.Kort
	kortStokklen := 5

	for i := len(hand) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		hand[i], hand[j] = hand[j], hand[i]
	}
	fmt.Println(hand)

	for i := 0; i < len(spillere); i++ {

		var playerHand []string

		for j := 0; j < kortStokklen; j++ {
			playerHand = append(playerHand, hand[j])
		}

		spDeck[spillere[i]] = playerHand
		spillerKort.spKort = spDeck
	}
	fmt.Println(spillerKort.spKort)
}