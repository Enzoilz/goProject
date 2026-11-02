package main

import "fmt"

func main() {
	slices1 := []int{12, 7, 4, 9, 16, 3, 10}
	fmt.Println(filterPair(slices1))
}

func filterPair(pair1 []int) []int {

	var result []int
	for _, v := range pair1 {
		if v%2 == 0 {
			result = append(result, v)
		}
	}

	return result
}
