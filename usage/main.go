package main

import (
	"fmt"
	"strings"
)

//go:generate gogenperllist Player

// Player is the content type of the perl list
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

	// NewPlayerList: create new list
	players := NewPlayerList(
		Player{"Messi", "Barcelona"},
		Player{"Ronaldo", "Real Madrid"},
		Player{"Neymar", "PSG"},
	)

	// Foreach: print each element
	players.Foreach(func(position int, p Player) {
		fmt.Printf("Number %d: %s plays for %s\n", position+1, p.Name, p.Club)
	})

	// Push: add new element to the end of the list
	players.Push(Player{"Müller", "FC Bayern"})
	players.Foreach(display)
	fmt.Printf("Position 2 is %s\n", players.ElementAt(1).Name)

	// Grep: filter element
	playersWithM := players.Grep(func(i int, p Player) bool {
		if strings.HasPrefix(p.Name, "M") {
			return true
		}
		return false
	})
	fmt.Println("Players with M")
	playersWithM.Foreach(display)

	// Map: modify elements
	playersUpcase := players.Map(func(i int, p Player) Player {
		p.Name = strings.ToUpper(p.Name)
		return p
	})
	fmt.Println("Players upcase")
	playersUpcase.Foreach(display)

	// Shift: get and remove first element
	first := players.Shift()
	fmt.Printf("First was  %s (%s)\n", first.Name, first.Club)
	players.Foreach(display)

	// Pop: get and remove last element
	last := players.Pop()
	fmt.Printf("Last was  %s (%s)\n", last.Name, last.Club)
	players.Foreach(display)

	// Shift: add new element at the beginning
	players.Unshift(Player{"Reus", "Dortmund"})
	fmt.Printf("Most famous %d players\n", players.Len())
	players.Foreach(display)

	// PopChecked: remove first element and get info if list is empty
	for last, ok := players.PopChecked(); ok; last, ok = players.PopChecked() {
		fmt.Printf("Last was  %s (%v)\n", last.Name, ok)
	}
	last, ok := players.PopChecked()
	fmt.Printf("\nLast was  %s (%v)\n", last.Name, ok)
}
