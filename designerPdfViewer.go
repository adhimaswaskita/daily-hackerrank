package main

import (
	"fmt"
	"strings"
)

func designerPdfViewer(h []int32, word string) int32 {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	alphabetDataset := strings.Split(alphabet, "")
	charDataset := strings.Split(word, "")
	var highest int32
	for _, char := range charDataset {
		for i, alphabetChar := range alphabetDataset {
			if char == alphabetChar && highest < h[i] {
				highest = h[i]
			}
		}
	}
	result := highest * int32(len(charDataset))
	return result
}

func main() {
	res := designerPdfViewer([]int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5}, "abc")
	fmt.Println(res)
}
