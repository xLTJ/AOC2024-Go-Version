package main

import (
	"fmt"
	"main/lib"
)

func main() {
	inputFile := "Input.txt"
	result, err := lib.CalculateDistance(inputFile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(result)
}
