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
		Speed: 240,
	}

	vehicle.Speed = 180

	fmt.Printf("votre véhicules : %s %v", vehicle.Brand, vehicle.Speed)
	fmt.Println(accelerate(vehicle, 40))
}

func accelerate(vehicle Vehicle, amount int) Vehicle {
	vehicle.Speed += amount

	return vehicle
}