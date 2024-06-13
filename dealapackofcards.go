package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	step := len(deck) / 4
	fmt.Printf("Player 1: %v", deck[0:step])
	fmt.Printf("Player 2: %v", deck[step:step*2])
	fmt.Printf("Player 3: %v", deck[step*2:step*3])
	fmt.Printf("Player 4: %v", deck[step*3:step*4])
}
