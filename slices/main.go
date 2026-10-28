package main

import "fmt"

func main() {

	age := []int{18, 21, 15, 30, 17}
	fmt.Println(calculeMoyenne(age))
}

func calculeMoyenne(ages []int) int {
	sum := 0

	for _, v := range ages {
		sum += v
	}

	return sum
}
