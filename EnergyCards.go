package main

import (
	_ "fmt"
)

var ()

//BasicEnergy ...
func BasicEnergy(F *Field, Type ...Element) *Field {
	Objective := PickPokemon(F.Player1.Bench)
	Objective.Energies.Cost[Type[0]]++
	return F
}

//SpecialEnergy ...
func SpecialEnergy(F *Field, Type ...Element) *Field {
	Objective := PickPokemon(F.Player1.Bench)
	for _, ty := range Type {
		Objective.Energies.Cost[ty]++
	}
	return F
}
