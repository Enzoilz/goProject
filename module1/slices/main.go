package main

import "fmt"

func main() {

	age := []float64{18, 21, 15, 30, 17}
	fmt.Println(calculeMoyenne(age))
}

func calculeMoyenne(ages []float64) float64 {
	if len(ages) == 0 {
		return 0.0
	}

	var sum float64

	for _, v := range ages {
		sum += v
	}

	return sum / float64(len(ages))
}
