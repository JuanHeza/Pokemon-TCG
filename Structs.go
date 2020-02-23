package main

import (
	"fmt"
	"image"
	"log"
	"math/rand"
	"reflect"
	"strconv"
	"time"
)

//Element is the pokemon type or the energy kind
type Element int
type rarity int
type status int

const (
	typeColorless Element = iota
	typeGrass
	typeFire
	typeWater
	typeLightning
	typePsychic
	typeFighting
	typeDarkness
	typeMetal
	typeFairy
	typeDragon

	common rarity = iota
	uncommon
	rare
	holofoilRare
	ultraRare
	secretRare

	//Normal is one of many status for a pokemon, Normal is the default
	Normal status = iota
	//Burned damage a pokemon each turn
	Burned
	//Poisoned Damage a pokemon each turn
	Poisoned
	//Asleep , the pokemon cant attack, flip a coin to awake
	Asleep
	//Confused , the pokemon may attack itself, flip a coin
	Confused
	//Paralyzed , the pokemon cant attack for one turn
	Paralyzed
	//SuperPoison damage a pokemon each turn
	SuperPoison

	//Head result of a coin
	Head = true
	//Tail result of a coin
	Tail = false
)

//Stages is a map of the actual (Gen 1) stages
var Stages = map[string]int{
	"Basic":   0,
	"Stage 1": 1,
	"Stage 2": 2,
}

//Rarities is a Map of rarities, used in coder & decoder
var Rarities = map[string]rarity{
	"Common":        common,
	"Uncommon":      uncommon,
	"Rare":          rare,
	"Holofoil Rare": holofoilRare,
	"Ultra Rare":    ultraRare,
	"Secret Rare":   secretRare,
}

//Elements is a map of elements used in the coder & decoder
var Elements = map[string]Element{
	"Colorless": typeColorless,
	"Grass":     typeGrass,
	"Fire":      typeFire,
	"Water":     typeWater,
	"Lightning": typeLightning,
	"Psychic":   typePsychic,
	"Fighting":  typeFighting,
	"Darkness":  typeDarkness,
	"Metal":     typeMetal,
	"Fairy":     typeFairy,
	"Dragon":    typeDragon,
}

//EnergyCost is a struct to store data about energy amounts and type
type EnergyCost struct {
	Cost map[Element]int
}

//------- DECK SECTION

//Deck is the deck data, Id, name, main types and Cards
type Deck struct {
	ID    string
	Title string
	Types []Element
	Cards [60]DeckCard
}

//DeckCard is the count of diferent cards in the deck
type DeckCard struct {
	ID     string
	Title  string
	Rarity rarity
	Count  int
}

//------- CARD SECTION

//CardBase is the meta data of a card
type CardBase struct {
	ID         string      `json:"id,omitempty"`
	Title      string      `json:"name,omitempty"`
	Imagecoord image.Point `json:"imagecoord,omitempty"`
	SetNumber  int         `json:"number,omitempty"`
	Series     string      `json:"series,omitempty"`
	Set        string      `json:"set,omitempty"`
	Setcode    string      `json:"setcode,omitempty"`
	Rarity     rarity      `json:"rarity,omitempty"`
	CardType   string      `json:"supertype,omitempty"`
	Info       interface{} `json:"info,omitempty"`
}

//CardPokemon stores the data of a pokeon type card
type CardPokemon struct {
	CardBase
	Hp         int
	MaxHp      int        `json:"hp,omitempty"`
	Stage      int        `json:"subtype,omitempty"`
	Type       Element    `json:"types,omitempty"`
	Slot1      Attacks    `json:"slot_1,omitempty"`
	Slot2      Attacks    `json:"slot_2,omitempty"`
	Energies   EnergyCost `json:"energies,omitempty"`
	Weaknesses EnergyCost `json:"weaknesses,omitempty"`
	Reatreat   EnergyCost `json:"reatreat,omitempty"`
	Resistence EnergyCost `json:"resistence,omitempty"`
	StatDamage status     `json:"stat_damage,omitempty"` // burned, poison & super poison
	StatMove   status     `json:"stat_move,omitempty"`   // paralyzed,sleep & confused
	Pokedex    int        `json:"nationalPokedexNumber,omitempty"`
	Preevolution string	  `json:"evolvesFrom,omitempty"`
}

//CardEnergy stores the data of a energy type card
type CardEnergy struct {
	CardBase
	Energy      EnergyCost
	Extra       string
	SubType     string `json:"sub_type,omitempty"` //special or basic
	Description string `json:"description,omitempty"`
}

//CardTrainer stores the data of a trainer type card
type CardTrainer struct {
	CardBase
	SubType     string `json:"sub_type,omitempty"` //special or basic
	Description string `json:"description,omitempty"`
	//Efects			map[string]Efect
}

func (PC *CardPokemon) printCard() {
	fmt.Println(PC)
}

func (TC *CardTrainer) printCard() {
	fmt.Println(TC)
}

func (EC *CardEnergy) printCard() {
	fmt.Println(EC)
}

//CardInterface is an interface for the 3 types of cards
type CardInterface interface {
	printCard()
}

//FlipCoin Flips a coin to get a value (Head = true | Tail = False)
func FlipCoin() bool {
	rand.Seed(time.Now().UnixNano())
	Res := rand.Intn(100) % 2
	if Res == 0 {
		return Head
	}
	return Tail
}

//ConsumeEnergy Consumes an energy attched to the pokemon  ***** MAY NOT USED *****
type ConsumeEnergy struct {
	Cost EnergyCost
}

//Attk is a data type to store all the attacks & pokepower
type Attk func(F *Field, Params ...int) *Field

//AttackMap stores the attacks via "Key = name | Value = func"
var AttackMap = map[string](Attk){}

//AttackMap[att].(field,Params)

//newPokemonCard Creates a new card using a map of interfaces
func newPokemonCard(dats map[string]interface{}) CardPokemon {
	hp, err := strconv.Atoi(fmt.Sprintf("%s", dats["hp"]))
	if err != nil {
		log.Println(err)
	}
	num, err := strconv.Atoi(fmt.Sprintf("%s", dats["number"]))
	if err != nil {
		log.Println(err)
	}
	Pdex, err := strconv.ParseFloat(fmt.Sprintf("%f", dats["nationalPokedexNumber"]), 64)
	if err != nil {
		log.Println(err)
	}
	var P = CardPokemon{
		CardBase: CardBase{
			ID:        fmt.Sprintf("%s", dats["id"]),
			Title:     fmt.Sprintf("%s", dats["name"]),
			Series:    fmt.Sprintf("%s", dats["series"]),
			Set:       fmt.Sprintf("%s", dats["set"]),
			Setcode:   fmt.Sprintf("%s", dats["setCode"]),
			Rarity:    Rarities[fmt.Sprintf("%s", dats["rarity"])],
			CardType:  fmt.Sprintf("%s", dats["supertype"]),
			SetNumber: num,
		},
		Hp:         hp,
		MaxHp:      hp,
		Stage:      Stages[fmt.Sprintf("%s", dats["subtype"])],
		Energies:   newEnergyCost(nil, -1),
		Weaknesses: newEnergyCost(dats["weaknesses"], 1),
		Reatreat:   newEnergyCost(dats["retreatCost"], 0),
		Resistence: newEnergyCost(dats["resistances"], 1),
		StatDamage: Normal,
		StatMove:   Normal,
		Pokedex:    int(Pdex),
	}

	if reflect.ValueOf(dats["types"]).Len() != 1 {
		log.Fatal()
	}
	P.Type = Elements[fmt.Sprintf("%s", reflect.ValueOf(dats["types"]).Index(0))]

	ev, ok := dats["evolvesFrom"]
	if ok{
		P.Preevolution = fmt.Sprintf("%s", ev)
	}

	ab, ok := dats["ability"]
	if !ok {
		log.Println("No Ability")
	}
	at, ok := dats["attacks"]
	if !ok {
		log.Println("No Attacks")
	}
	P.Slot1, P.Slot2 = newAttacks(at, ab)
	return P
}

func newEnergyCost(in interface{}, val int) EnergyCost {
	if in != nil {
		switch val {
		case 0:
			var mp = make(map[Element]int)
			ln := reflect.ValueOf(in).Len()
			for i := 0; i < ln; i++ {
				in := reflect.ValueOf(in).Index(i)
				mp[Elements[fmt.Sprintf("%s", in)]]++
			}
			return EnergyCost{Cost: mp}
		case 1:
			Y := reflect.ValueOf(in).Index(0)
			l := Y.Interface().(map[string]interface{})
			if fmt.Sprintf("%s", l["value"]) == "×2" {
				return EnergyCost{
					Cost: map[Element]int{
						Elements[fmt.Sprintf("%s", l["type"])]: 2,
					},
				}
			}
			d, err := strconv.Atoi(fmt.Sprintf("%s", l["value"]))
			if err != nil {
				log.Println(err)
			}
			return EnergyCost{
				Cost: map[Element]int{
					Elements[fmt.Sprintf("%s", l["type"])]: d,
				},
			}
		}
	}
	return EnergyCost{}
}

func newAttacks(at interface{}, ab ...interface{}) (Attacks, Attacks) {
	var slot1, slot2 = Attacks{}, Attacks{}
	if ab[0] != nil {
		Y := reflect.ValueOf(ab).Index(0).Interface().(map[string]interface{})
		slot1 = Attacks{
			Name:        fmt.Sprintf("%s", Y["name"]),
			Description: fmt.Sprintf("%s", Y["text"]),
			Ability:     true,
			Active:      true,
			//Do:,
		}
		l := reflect.ValueOf(at).Index(0).Interface().(map[string]interface{})
		slot2 = slotFillingAttack(l)
	} else {
		l := reflect.ValueOf(at).Index(0).Interface().(map[string]interface{})
		slot1 = slotFillingAttack(l)
		if reflect.ValueOf(at).Len() > 1 {
			l := reflect.ValueOf(at).Index(0).Interface().(map[string]interface{})
			slot2 = slotFillingAttack(l)
		}

	}
	return slot1, slot2
}

func slotFillingAttack(l map[string]interface{}) Attacks {
	var slot = Attacks{
		Name:        fmt.Sprintf("%s", l["name"]),
		Description: fmt.Sprintf("%s", l["text"]),
		Ability:     false,
		Active:      true,
	}

	is := fmt.Sprintf("%s", l["damage"])
	if is != "" {
		P, err := strconv.Atoi(is)
		slot.Params = append(slot.Params, P)
		if err != nil {
			log.Println(err)
		}
	}

	ln := reflect.ValueOf(l["cost"]).Len()
	var mp = make(map[Element]int)
	for i := 0; i < ln; i++ {
		in := reflect.ValueOf(l["cost"]).Index(i)
		mp[Elements[fmt.Sprintf("%s", in)]]++
	}
	slot.Cost = EnergyCost{
		Cost: mp,
	}
	return slot
}

func newEnergyCard(dats map[string]interface{}) CardEnergy {
	num, err := strconv.ParseFloat(fmt.Sprintf("%f", dats["number"]), 64)
	if err != nil {
		log.Println(err)
	}
	var E = CardEnergy{
		CardBase: CardBase{
			ID:        fmt.Sprintf("%s", dats["id"]),
			Title:     fmt.Sprintf("%s", dats["name"]),
			Series:    fmt.Sprintf("%s", dats["series"]),
			Set:       fmt.Sprintf("%s", dats["set"]),
			Setcode:   fmt.Sprintf("%s", dats["setCode"]),
			Rarity:    Rarities[fmt.Sprintf("%s", dats["rarity"])],
			CardType:  fmt.Sprintf("%s", dats["supertype"]),
			SetNumber: int(num),
		},
		SubType: fmt.Sprintf("%s", dats["subtype"]),
	}
	if E.SubType != "Basic" {
		E.Description = fmt.Sprintf("%s", dats["text"])
	}
	return E
}

func newTrainerCard(dats map[string]interface{}) CardTrainer {
	num, err := strconv.ParseFloat(fmt.Sprintf("%f", dats["number"]), 64)
	if err != nil {
		log.Println(err)
	}
	var E = CardTrainer{
		CardBase: CardBase{
			ID:        fmt.Sprintf("%s", dats["id"]),
			Title:     fmt.Sprintf("%s", dats["name"]),
			Series:    fmt.Sprintf("%s", dats["series"]),
			Set:       fmt.Sprintf("%s", dats["set"]),
			Setcode:   fmt.Sprintf("%s", dats["setCode"]),
			Rarity:    Rarities[fmt.Sprintf("%s", dats["rarity"])],
			CardType:  fmt.Sprintf("%s", dats["supertype"]),
			SetNumber: int(num),
		},
		Description: fmt.Sprintf("%s", dats["text"]),
		SubType:     fmt.Sprintf("%s", dats["subtype"]),
	}
	return E
}

/*
	evolvesFrom:Kadabra
*/
