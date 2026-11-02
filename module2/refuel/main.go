package main

import "fmt"

type Vehicle struct {
    Brand string
    Fuel  int
}

func main() {
    vehicle := Vehicle{
        Brand: "BMW",
        Fuel: 70,
    }

    refuel(&vehicle, 20)

    fmt.Println(vehicle.Fuel)
}

func refuel(vehicle *Vehicle, amount int)  {
	if amount <= 0 {
		return
	}

	if amount + vehicle.Fuel >= 100 {
		vehicle.Fuel = 100
		fmt.Println("*Le véhicule ne peux pas dépasser 100 de carburant")
		return
	}

	vehicle.Fuel += amount
}