package main

import "fmt"

type Vehicle struct {
    Brand string
    Speed int
    Fuel  int
}

func main() {
    vehicle := Vehicle{
        Brand: "BMW",
        Speed: 100,
        Fuel: 80,
    }

    drive(&vehicle, 40)

    fmt.Println(vehicle.Speed)
    fmt.Println(vehicle.Fuel)
}

func drive(vehicle *Vehicle, distance int,)  {
	vehicle.Speed += distance
	vehicle.Fuel -= distance /2
}