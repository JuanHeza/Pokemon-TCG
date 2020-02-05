package main
import(
    "log"
    "math/rand"
    //"time"
)
var(
	Attaking, Defending *Side
)

type Player struct{
	Hand 	[]Card 
	Deck	[]Card
	Trash	[]Card
	Prize	[]Card
	Side	*Side
}

type Side struct{
	Main 	Pokemon
	Bench1	Pokemon
	Bench2 	Pokemon
	Bench3 	Pokemon
	Bench4 	Pokemon
	Bench5 	Pokemon
}

type Field struct{
	Player1 *Side
	Player2 *Side
	stadium Card
}

type Card struct{
	Name 	string	// title
	Desc	string	// description
	Set		string	// card set
	Number	int		// number in set
	Kind	string	// pokemon, trainer, energy
	Stage	string	// in case of pokemon card
}

func Battle(Pl1 string, Pl2 string){
	var BattleField Field
	var Player1, Player2 Player
	Attaking, Defending = Player1.Side, Player2.Side
	Player1.Init(Pl1)
	Player2.Init(Pl2)
	Player1.Side = BattleField.Player1
	Player2.Side = BattleField.Player2

}

func (p *Player) Init( Pl string){
	p.GetDeck(Pl)
	p.ShuffleDeck()
	p.GetHand()
}

func (p *Player) GetDeck(Pl string){
	/*
		READ DATA AND GET THE DECK OF PL
		p.Deck = that
	*/
}

func (p *Player)  ShuffleDeck(){
	//Seed := rand.New(rand.NewSource(time.Now().UnixNano()))
	rand.Shuffle(len(p.Deck), func(i, j int) { p.Deck[i], p.Deck[j] = p.Deck[j], p.Deck[i] })
}

func (p *Player) GetHand(){
	var a,b = 0,5
	var pokecard = false
	for !pokecard{
		p.Hand = p.Deck[a:b]
		for _,x := range p.Hand{
			if x.Stage == "Basic"{
				pokecard = true
			}
		}
		log.Println("No Basic Pokemon")
		a, b = b, b+5
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