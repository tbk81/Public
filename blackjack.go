package main

import (
	"fmt"
	"github.com/tbk81/100-DAYS-OF-CODE-GO/0-tools/randoGen"
)

// var deck = make([]string,52)
var deck = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10}

var playerHand []int
var computerHand []int

func dealHand(i int, d []int) []int {
	return randoGen.Rando(i, d)
}

func drawCard() []int {
	return randoGen.Rando(1, deck)
}

func checkHand(h []int) bool {
	var t int
	for _, v := range h {
		t += v
	}
	fmt.Println(t)
	if t > 21 {
		return true
	} else {
		return false
	}
}

func main() {
	var usrChoice string
	var hit string
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
			fmt.Printf("Your cards: %v\nComputer's first card: %v\n", playerHand, computerHand[0])
		Loop2:
			for {
				fmt.Print("'h' to hit or 'p' to pass: ")
				fmt.Scanln(&hit)
				switch hit {
				case "h":
					playerHand = append(playerHand, drawCard()[0])
					fmt.Println(playerHand)
					fmt.Println(checkHand(playerHand))
				case "p":
					break Loop2
				}
			}
		case "n":
			break Loop1
		}
	}
}

package randoGen

import (
	"math/rand"
	"time"
)

// const symset = "!?@#$%^&*"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func SliceWithRandoset(length int, xi []int) []int {
	b := make([]int, length)
	for i := range b {
		b[i] = xi[seededRand.Intn(len(xi))]
	}
	return b
}

func Rando(length int, xi []int) []int {
	return SliceWithRandoset(length, xi)
}
