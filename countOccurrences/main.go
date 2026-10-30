package main

import "fmt"

func main() {

	numbers := []int{2, 5, 2, 8, 5, 2, 9}
	fmt.Println(countOccurrences(numbers))
}

func countOccurrences(number []int) map[int]int {
	num1 := make(map[int]int)

	for _, v := range number {
		num1[v]++
	}

	return num1
}
