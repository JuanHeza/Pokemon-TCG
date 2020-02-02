package main

import (
   
)

type Elements int

const(
	Water Elements = iota // Water & Ice			// Blue
	Nature 				// Bug & Grass			// Green
	Fire 				// Fire					// Red
	Lightning 			// Electric				// Yellow
	Psychic 			// Psychic & Ghost		// Purpule
	Fighting			// Fight, Rock & Ground	// Orange?Brown
	Darkness 			// Dark & Poison		// Black
	Metal 				// Steel				// Silver
	Fairy	 			// Fairy				// Pink
	Colorless 			// Normal & Flying		// White
	Dragon	 			// Dragon				// Gold
)
//Pokemon Card Info
/*
type Pokemon struct{
	Nombre string
	HP int
	Type []int
	Stage stage
	Weakness int
	Resistance int
	RetreatCost int
	Ability ability
	Attack []attack
	Rarity int
	Collection string
	Serial string
	Pokedex int
	Art string // image.Image
}
*/
type stage struct{
	stage int // basic = 0, stage 1 = 1, stage 2 = 2
	Prev int // # of the pokemon in the pokedex
}

//New create a Pokemon Card
/* func (PC *Pokemon) New(){
	fmt.Println("Name:")
	fmt.Scanf("%s\n",&PC.Nombre)
	fmt.Println("HP:")
	fmt.Scanf("%d\n",&PC.HP)
	fmt.Println("Weakness:")
	fmt.Scanf("%d\n",&PC.Weakness)
	fmt.Println("Rarity:")
	fmt.Scanf("%d\n",&PC.Rarity)
	fmt.Println("Pokedex:")
	fmt.Scanf("%d\n",&PC.Pokedex)
} */