package main

import "fmt"

type Player struct {
	Name string
	HP int
}

func main() {
	player := &Player{
		Name: "Enzo",
		HP: 50,
	}

	player.heal(30)

	fmt.Println(player)
}

func (p *Player) heal(amount int) {
	if amount <= 0 {
		fmt.Println("Le heal ne peux pas contenir un nombre négatif!")
		return
	}

	if p.HP + amount >= 100 {
		p.HP = 100
		fmt.Println("Le joueur ne peux pas dépasser 100 hp")
		return
	}
	
	fmt.Println("Le joueur vient de récupéré :", amount, "HP")
	p.HP += amount
}