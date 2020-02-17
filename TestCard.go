package main

import (
	"image"
)

var (
	//TestPokemon is a card for tests funcions
	TestPokemon = CardBase{
		ID:         "base1-17",
		Title:      "Beedrill",
		Imagecoord: image.Point{90, 55},
		SetNumber:  17,
		Series:     "Base",
		Set:        "Base",
		Setcode:    "base1",
		Rarity:     rare,
		//	CardType		Cards
		Info: CardPokemon{
			Hp:    80,
			MaxHp: 80,
			Energies: EnergyCost{
				Cost: make(map[Element]int),
			},
			Stage: 2,
			Type:  typeGrass,
			Slot1: Attacks{
				Name:        "Twineedle",
				Cost:        EnergyCost{Cost: map[Element]int{typeColorless: 3}},
				Description: "Flip 2 coins. This attack does 30 damage times the number of heads.",
				Active:      true,
				Do:          Twineedle,
			},
			Slot2: Attacks{
				Name:        "Poison Sting",
				Cost:        EnergyCost{Cost: map[Element]int{typeGrass: 3}},
				Description: "Flip a coin. If heads, the Defending Pokémon is now Poisoned.",
				Active:      true,
				Params:      []int{22, 5},
				Do:          PoisonSting,
			},
			Weaknesses: EnergyCost{Cost: map[Element]int{typeFire: 2}},
			Reatreat:   EnergyCost{},
			Resistence: EnergyCost{Cost: map[Element]int{typeFighting: -30}},
			Pokedex:    15,
		},
	}
)
