package main

import "fmt"

func main() {

	age := 15

	if age >= 18 {
		fmt.Println("Vous avez plus de 18 ans, vous êtes majeur")
	} else {
		fmt.Println("Vous avez moins de 18 ans, vous êtes mineur")
	}
}
