package main

import (
	"fmt"
)

func main() {

	card := "three"
	value := ParseCard(card)
	fmt.Println(value)
}

func ParseCard(card string) int {
	switch {
	case card == "ace":
		return 11
	case card == "king" || card == "queen" || card == "jack" || card == "ten":
		return 10
	case card == "nine":
		return 9
	case card == "eight":
		return 8
	case card == "seven":
		return 7
	case card == "six":
		return 6
	case card == "five":
		return 5
	case card == "four":
		return 4
	case card == "three":
		return 3
	case card == "two":
		return 2
	case card == "one":
		return 1
	default:
		return 0
	}
	panic("Please implement the ParseCard function")
}
