package main

import "fmt"

func breakingRecords(scores []int32) (response []int32) {
	lowestScore := scores[0]
	highestScore := scores[0]
	var low, hi int32
	for _, score := range scores {
		if score < lowestScore {
			lowestScore = score
			low++
		} else if score > highestScore {
			highestScore = score
			hi++
		}
	}
	response = []int32{hi, low}
	return
}

func main() {
	record := breakingRecords([]int32{10, 5, 20, 20, 4, 5, 2, 25, 1})
	fmt.Println(record)
}
