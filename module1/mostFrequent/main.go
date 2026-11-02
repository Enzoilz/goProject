package main

import "fmt"

func main() {

	numbers := []int{2, 5, 2, 8, 5, 2, 9}
	fmt.Println(mostFrequent(numbers))
}

func mostFrequent(number []int) int {
	num1 := make(map[int]int)

	var maxCount int
	var mostNumber int

	for _, v := range number {
		num1[v]++
	}

	for i, v := range num1 {

		if v >= maxCount {
			maxCount = v
			mostNumber = i
		}
	}

	return mostNumber
}
