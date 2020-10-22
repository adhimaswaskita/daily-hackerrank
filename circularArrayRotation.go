package main

import (
	"fmt"
	"strconv"
	"strings"
)

func rotator(payload string, arrValue string) ([]int, error) {
	payloadDataSet := strings.Split(payload, "")
	arrValueDataset := strings.Split(arrValue, "")

	lengthPayload, _ := strconv.Atoi(payloadDataSet[0])
	rotationPayload, _ := strconv.Atoi(payloadDataSet[1])
	queryPayload, _ := strconv.Atoi(payloadDataSet[2])

	convertedArrValueDataset := make([]int, lengthPayload)
	for i := 0; i < len(arrValueDataset); i++ {
		convertedArrValue, err := strconv.Atoi(arrValueDataset[i])
		if err != nil {
			return nil, err
		}
		convertedArrValueDataset[i] = convertedArrValue
	}

	var responseArray []int
	for i := 0; i < rotationPayload; i++ {
		for x := 0; x < (len(convertedArrValueDataset) - 1); x++ {
			responseArray = append(responseArray, convertedArrValueDataset[x])
		}
		responseArray = append([]int{convertedArrValueDataset[len(convertedArrValueDataset)-1]}, responseArray...)
		convertedArrValueDataset = responseArray[:len(convertedArrValueDataset)]
	}

	indexToShow := make([]int, queryPayload)
	for i := 0; i < queryPayload; i++ {
		fmt.Scanln(&indexToShow[i])
	}

	var newArray []int
	for i := 0; i < len(indexToShow); i++ {
		for y := 0; y < len(responseArray); y++ {
			if indexToShow[i] == y {
				newArray = append(newArray, responseArray[y])
			}
		}
	}

	return newArray, nil
}

func main() {
	var payload string
	var arrayData string
	fmt.Scanln(&payload)
	fmt.Scanln(&arrayData)

	rotatedData, _ := rotator(payload, arrayData)
	for _, data := range rotatedData {
		fmt.Println(data)
	}
}
