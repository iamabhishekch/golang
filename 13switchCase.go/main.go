package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("switch cases in go lang ")

	rand.New(rand.NewSource(time.Now().UnixNano())) // understand this deeply
	diceNumber := rand.Intn(6) + 1
	fmt.Println("value of dice: ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("the value is 1 now you can open")
	case 2:
		fmt.Println("the value is 2 now you can move 2 steps")
	case 3:
		fmt.Println("the value is 3 now you can move 3 steps")
	case 4:
		fmt.Println("the value is 4 now you can move 4 steps")
		fallthrough // i.e., execute the next case block even if it doesn't match
	case 5:
		fmt.Println("the value is 5 now you can move 5 steps")
		fallthrough //  Fall through to the next case
	case 6:
		fmt.Println("the value is 6 now you can move 6 steps and roll dice again")
	default:
		fmt.Println("you are not allowed to play this game")

	}

}
