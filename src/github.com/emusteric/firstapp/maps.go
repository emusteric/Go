package main

import (
	"fmt"
)

func main() {
	statePopulations := map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}
	fmt.Println(statePopulations)
	fmt.Println(statePopulations["Texas"])
	statePopulations["Georgia"] = 10310371
	fmt.Println(statePopulations)
	delete(statePopulations, "Georgia")
	fmt.Println(statePopulations)

	for k, v := range statePopulations {
		fmt.Println(k, v)
	}

}
