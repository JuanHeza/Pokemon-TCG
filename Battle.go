package main

import (
	"log"
	"math/rand"
	//"time"
)

var (
	//Adversary is a temporal variable, not using
	Adversary            Player
	attacking, defending *Side
)

//Player data struct
type Player struct {
	Hand  []Card
	Deck  []Card
	Trash []Card
	Prize []Card
	side  *Side
}

//Side is the field side of the player
type Side struct {
	Main      *CardPokemon
	Bench     [5]*CardPokemon
	CanEvolve [6]bool
}

//Field is the total play area, both sides
type Field struct {
	Player1 *Side
	Player2 *Side
	stadium Card
}

//Card is the struct of a card in the deck
type Card struct {
	Name   string // title
	Desc   string // description
	Set    string // card set
	Number int    // number in set
	Kind   string // pokemon, trainer, energy
	Stage  string // in case of pokemon card
}

//lastAttack done results
type lasAttack struct {
	Name       string
	Damage     int
	StatMove   status
	StatDamage status
}

//LastAttack variable
var LastAttack lasAttack

//Battle funcion, here start the game
func Battle(Pl1 string, Pl2 string) {
	var BattleField Field
	var Player1, Player2 Player
	attacking, defending = Player1.side, Player2.side
	Player1.Init(Pl1)
	Player2.Init(Pl2)
	Adversary = Player2
	Player1.side = BattleField.Player1
	Player2.side = BattleField.Player2
}

//Init is the player initializer in the game
func (p *Player) Init(Pl string) {
	p.GetDeck(Pl)
	p.ShuffleDeck()
	p.GetHand()
}

//GetDeck search the player in a list and return the respective deck
func (p *Player) GetDeck(Pl string) {
	/*
		READ DATA AND GET THE DECK OF PL
		p.Deck = that
	*/
}

//ShuffleDeck ... well shuffle the deck
func (p *Player) ShuffleDeck() {
	//Seed := rand.New(rand.NewSource(time.Now().UnixNano()))
	rand.Shuffle(len(p.Deck), func(i, j int) { p.Deck[i], p.Deck[j] = p.Deck[j], p.Deck[i] })
}

//GetHand gives a player hand
func (p *Player) GetHand() {
	var a, b = 0, 5
	var pokecard = false
	for !pokecard {
		p.Hand = p.Deck[a:b]
		for _, x := range p.Hand {
			if x.Stage == "Basic" {
				pokecard = true
			}
		}
		log.Println("No Basic Pokemon")
		a, b = b, b+5
	}
}

//BattleTest is the thest of battle functions
func BattleTest() {
	var def CardPokemon
	var att CardPokemon
	pokeside := FlipCoin()
	if pokeside {
		def = Sets["Base"]["base1-1"].(CardPokemon)
		att = Sets["Base"]["base1-3"].(CardPokemon)
	} else {
		def = Sets["Base"]["base1-3"].(CardPokemon)
		att = Sets["Base"]["base1-1"].(CardPokemon)
	}
	drawSetPokemon(win, def, "P2-Main")
	drawSetPokemon(win, att, "P1-Main")
	/*
	 */
}

func playerTurn() {
	var Energyused, Attacked  = false, false
	//var PlacePokemon int
	//var Evolvepokemon int
	//var Reemplace int
	for !Attacked {
		if Energyused{
			
		}
	}

}

//EvolvePokemon evolves a pokemon to the next stage
func (P *CardPokemon) EvolvePokemon(EP *CardPokemon) {
	if EP.Preevolution == P.Title {
		EP.Hp -= (P.MaxHp - P.Hp)
	}

}

/*
forums.pokemontcg.com/topic/50237-original-decks-1999-2001/

haymaker
*	4	Hitmonchan
*	4	Electabuzz
*	4	Scyther
*	4	Plus Power
*	4	Energy Removal
*	4	Super Energy Removal
*	3	Gust of Wind
*	4	Oak
*	3	Bill
*	1	Lass
*	2	Item Finder
*	2	Computer Search
*	4	DCE
*	9	Fighting Energy
*	8 	Lighting Energy
*
Turbo Rain Dance
*	4	Squirtle
*	4	Blastoise
*	4	Pokemon Breeder
*	4	Oak
*	4	Bill
*	3	Pokemon Trader
*	4	Item Finder
*	4	Computer Search
*	2	Gust of Wind
*	3	Goop Gas Attack
*	2	Plus Power
*	2	Defender
*	1	Full Heal
*	2	Super Energy Removal
*	17	Water Energy
*
Fightmare
*	3	Hitmonlee
*	3	Hitmonchan
*	3	Gastly - Fossil
*	3	Haunter - Fossil
*	2	Gengar
*	2	Mankey
*	2	Oak
*	4	Bill
*	1	Pokemon Trader
*	4	Energy Removal
*	2	Computer Search
*	2	Gust of Wind
*	2	Scoop Up
*	2	Plus Power
*	1	Energy Retrival
*	3	Super Energy Removal
*	12	Fighting Energy
*	9	Psychic Energy
*
*/
