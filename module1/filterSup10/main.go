package main

import "fmt"

func main() {
	slices := []int{12, 7, 4, 19, 10, 15}
	slices2 := []int{}
	fmt.Println(filterSup10(slices))
	fmt.Println(filterSup10(slices2))
}

func filterSup10(slice []int) []int {
	var some []int
	for _, v := range slice {
		if v > 10 {
			some = append(some, v)
		}
	}
	return some
}
