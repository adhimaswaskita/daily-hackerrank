package main

import (
	"fmt"
	"strings"
)

func camelCaseCounter(message string) (total int16) {
	runes := []rune(message)
	for i, rune := range runes {
		if string(rune) == strings.ToUpper(string(rune)) || i == 0 {
			total += 1
		}
	}
	return
}

func main() {
	var message string
	fmt.Print("Message \t: ")
	fmt.Scanf("%s \n", &message)
	fmt.Printf("Total \t\t: %v \n", camelCaseCounter(message))

}