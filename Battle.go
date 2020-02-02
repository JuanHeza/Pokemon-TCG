package main
import(
    "log"
    "math/rand"
    //"time"
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