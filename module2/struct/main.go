package main

import "fmt"

type Vehicle struct {
	Brand string
	Model string
	Speed int 
}

func main()  {
	

	vehicle := Vehicle{
		Brand: "BMW",
		Model: "M4",
		Speed: 240
	}

	fmt.Println(vehicle.Brand)
}

