package main

import (
	"fmt"

	"github.com/tbk81/100-DAYS-OF-CODE-GO/00-tools/randoGen"
)

var deck = []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 11, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 11, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 11, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 11}

func dealHand(i int, d []int) []int {
	return randoGen.Rando(i, d)
}

func drawCard() []int {
	return randoGen.Rando(1, deck)
}

func score(h []int) int {
	var t int
	for _, v := range h {
		t += v
	}
	return t
}

func checkHand(h []int) bool {
	t := score(h)
	if t > 21 {
		return true
	} else {
		return false
	}
}

func playComp(h []int) []int {
	if checkHand(h) == false && score(h) < 17 {
		h = append(h, drawCard()[0])
	}
	return h
}

func main() {
	var playerHand, computerHand []int
	var usrChoice, hit string
	fmt.Println("Welcome to blackjack!")
Loop1:
	for {
		fmt.Print("Do you want to play a game of blackjack? Type 'y' or 'n': ")
		fmt.Scanln(&usrChoice)
		switch usrChoice {
		case "y":
			fmt.Println("Start playing game")
			playerHand = dealHand(2, deck)
			computerHand = dealHand(2, deck)
			fmt.Printf("Your cards: %v\tCurrent total: %v\nComputer's first card: %v\n", playerHand, score(playerHand), computerHand[0])
		Loop2:
			for {
				fmt.Print("'h' to hit or 'p' to pass: ")
				fmt.Scanln(&hit)
				switch hit {
				case "h":
					playerHand = append(playerHand, drawCard()[0])
					if playerHand[len(playerHand)-1] == 11 && checkHand(playerHand) {
						playerHand[len(playerHand)-1] = 1
					}
					fmt.Printf("Your hand: %v\tCurrent total: %v\n", playerHand, score(playerHand))
					fmt.Println("Computer's first card:", computerHand[0])
					if checkHand(playerHand) {
						fmt.Println("BUST! Computer wins")
						break Loop2
					}
				case "p":
					fmt.Println(playComp(computerHand))
					// clearScreen.Clear()
					break Loop2
				}
			}
		case "n":
			break Loop1
		}
	}
}
