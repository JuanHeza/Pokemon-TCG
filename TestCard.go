package main

import(
	"image"
)
var(
	TestPokemon = Card_Pokemon{
		ID:				"base1-17",
		Title:			"Beedrill",
		Imagecoord:		image.Point{90,55},
		SetNumber:		17,
		series:			"Base",
		Set:			"Base",
		Setcode:		"base1",
		Rarity:			Rare,
	//	CardType		Cards
		Stage:			"Stage 2",
		Type:			Grass,	
		Slot1:			Attack{
			Title:			"Twineedle",
			Cost:			EnergyCost{Cost: map[Element]int{Colorless:3}},
			Description:	"Flip 2 coins. This attack does 30 damage times the number of heads.",
			BaseDamage:		map[string]int{"X":30},//[.]10 -> 10 damage , [+]10 -> +10 damage ...
			Efects:			map[string]Efect{"Coin":Coin_Add{
				Coins: 		2,
				BaseDamage: 30,
				Operator:	"X",
				Additional:	0,
			},},
		},
		Slot2:			Attack{
			Title: 			"Poison Sting",
			Cost:			EnergyCost{Cost: map[Element]int{Grass:3}},
			Description: 	"Flip a coin. If heads, the Defending Pokémon is now Poisoned.",
			BaseDamage:		map[string]int{".":40},
			Efects:			map[string]Efect{"Coin":Coin_Stat{
				Coins: 			1,
				Head: 			Poisoned,
				Objective_Head: Defending.Main,
			},},
		},
		Weaknesses:		EnergyCost{Cost: map[Element]int{Fire:2}},
		Reatreat:		EnergyCost{},
		Resistence:		EnergyCost{Cost: map[Element]int{Fighting:-30}},
		Pokedex:		15,
	}

	TestTrainer = Card_Trainer{
		ID:				"base1-95",
		Title:			"Switch",
		Imagecoord:		image.Point{99,88},
		SetNumber:		95,
		series:			"Base",
		Set:			"Base",
		Setcode:		"Base",
		Rarity:			Common,
		//Energy			EnergyCost
		SubType:		"",	//special or basic
		Description:	"Switch 1 of your Benched Pokémon with your Active Pokémon.",
		//Efects:			map[string]Efect{"Switch":Switch{
		//	Main: Attacking.Main,
		//	Bench: Attaking,
		//},},
		//CardType		Cards
	}
)