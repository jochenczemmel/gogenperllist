package main

import (
	"fmt"
	"math/big"
)

//go:generate gogenperllist -import "math/big" Int

// show simple use cases
func demoMath() {

	fmt.Println("--- demoMath ---")

	// NewPlayerList: create new list
	numbers := NewIntList(
		*big.NewInt(1),
		*big.NewInt(2),
		*big.NewInt(3),
		*big.NewInt(4),
		*big.NewInt(5),
	)

	// Foreach: print each element
	numbers.Foreach(func(position int, n big.Int) {
		fmt.Printf("Number %d: is %v\n", position+1, n.Text(10))
	})

	gt3 := numbers.Any(func(i int, n big.Int) bool {
		return n.Cmp(big.NewInt(3)) == 1
	})
	fmt.Println("gibt es elemente > 3? ", gt3)
	gt3 = numbers.All(func(i int, n big.Int) bool {
		return n.Cmp(big.NewInt(3)) == 1
	})
	fmt.Println("Alle elemente > 3? ", gt3)

	fmt.Println("Pop:", numbers.Pop())
	fmt.Println("Pop:", numbers.Pop())
	fmt.Println("Shift:", numbers.Shift())
	fmt.Println("Shift:", numbers.Shift())
	fmt.Println("Pop:", numbers.Pop())
	fmt.Println("Pop:", numbers.Pop())
}
