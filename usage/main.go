package main

import (
	"fmt"
)

// content type of perl list
//go:generate gogenperllist Player
type Player struct {
	Name string
	Club string
}

// demo function for Foreach() method
func display(position int, p Player) {
	fmt.Printf("Number %d: %s plays for %s\n", position+1, p.Name, p.Club)
}

// show simple use cases
func main() {

	players := NewPlayerList(
		Player{"Messi", "Barcelona"},
		Player{"Ronaldo", "Real Madrid"},
		Player{"Neymar", "PSG"},
	)

	players.Foreach(func(position int, p Player) {
		fmt.Printf("Number %d: %s plays for %s\n", position+1, p.Name, p.Club)
	})

	players.Push(Player{"Müller", "FC Bayern"})
	players.Foreach(display)
	fmt.Printf("Position 2 is %s\n", players.ElementAt(1).Name)

	first := players.Shift()
	fmt.Printf("First was  %s (%s)\n", first.Name, first.Club)
	players.Foreach(display)

	last := players.Pop()
	fmt.Printf("Last was  %s (%s)\n", last.Name, last.Club)
	players.Foreach(display)

	players.Unshift(Player{"Reus", "Dortmund"})
	fmt.Printf("Most famous %d players\n", players.Len())
	players.Foreach(display)

	last, ok := players.PopChecked()
	fmt.Printf("Last was  %s (%v)\n", last.Name, ok)
	last, ok = players.PopChecked()
	fmt.Printf("Last was  %s (%v)\n", last.Name, ok)
	last, ok = players.PopChecked()
	fmt.Printf("Last was  %s (%v)\n", last.Name, ok)
	last, ok = players.PopChecked()
	fmt.Printf("Last was  %s (%v)\n", last.Name, ok)
}
