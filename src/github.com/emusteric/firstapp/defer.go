package main

import (
	"fmt"
)

func main() {
	fmt.Println("start")
	fmt.Println("middle")
	fmt.Println("end")

	fmt.Println("start")
	defer fmt.Println("middle")
	fmt.Println("end")
}
