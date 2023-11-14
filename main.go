package main

import (
	"fmt"
	"math/rand"
)

type Deck struct {
	Kort []string
}

type Gamers struct {
	Antall int
	navn string
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
	shuffleNgive(deck,gamers,spillerkort)
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
	var navn string
	var navnArr []string

	fmt.Print("Hvor mange spillere er det? ")
	fmt.Scan(&Antall)

	switch {
		case Antall <= 1:
			fmt.Println("Vennligst få flere til å være med")
		case Antall > 8:
			fmt.Println("For mange spillere")
		default:
			Gamers.Antall = Antall
	}

	
	for  i := 0;  i < Antall;  i++ {
		fmt.Print("Hva heter de? ")
		fmt.Scan(&navn)
		navnArr = append(navnArr, navn)
	}

	Gamers.navnArr = navnArr
	Gamers.navn = navn

	fmt.Println(navnArr)
}

func shuffleNgive(deck *Deck,Gamers *Gamers, spillerKort *spillerKort) {

	spillere := Gamers.navnArr
	spillKort := spillerKort.spKort
	hand := deck.Kort
	kortStokklen := 5

	for i := len(hand) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		hand[i], hand[j] = hand[j], hand[i]
	}
	fmt.Println(hand)

	
}