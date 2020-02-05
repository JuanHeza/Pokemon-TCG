package main

import(
	"image"
	"fmt"
    "math/rand"
    "time"
)

type Element int
type rarity int 
type status int
const(
	Colorless 	Element = iota
	Grass 		
	Fire 		
	Water			
	Lightning	
	Psychic		
	Fighting	
	Darkness	
	Metal		
	Fairy		
	Dragon		

	Common 			rarity = iota
	Uncommon
	Rare
	Holofoil_Rare
	Ultra_Rare
	Secret_Rare

	Burned 		status = iota
	Poisoned
	Asleep
	Confused
	Paralized
)

type EnergyCost struct{
	Cost map[Element]int
}

//------- DECK SECTION

type Deck struct{
	ID		string
	Title	string
	Types	[]Element
	Cards	[60]Deck_Card
}

type Deck_Card struct{
	ID		string
	Title	string
	Rarity	rarity
	Count	int
}

//------- CARD SECTION

type Card_Base struct{
	ID				string
	Title			string
	Imagecoord		image.Point
	SetNumber		int
	series			string
	Set				string
	Setcode			string
	Rarity			rarity
}

type Card_Pokemon struct{
	ID				string
	Title			string
	Imagecoord		image.Point
	SetNumber		int
	series			string
	Set				string
	Setcode			string
	Rarity			rarity
	//CardType		Cards
	Stage			string
	Type			Element	
	Slot1			CardSlot
	Slot2			CardSlot
	Weaknesses		EnergyCost
	Reatreat		EnergyCost
	Resistence		EnergyCost
	Pokedex			int
}

type Card_Energy struct{
	ID				string
	Title			string
	Imagecoord		image.Point
	SetNumber		int
	series			string
	Set				string
	Setcode			string
	Rarity			rarity
	Energy			EnergyCost
	SubType			string	//special or basic
	Description		string
	//CardType		Cards
}

type Card_Trainer struct{
	ID				string
	Title			string
	Imagecoord		image.Point
	SetNumber		int
	series			string
	Set				string
	Setcode			string
	Rarity			rarity
	//Energy			EnergyCost
	SubType			string	//special or basic
	Description		string
	Efects			map[string]Efect
	//CardType		Cards
}

type Cards interface{
	Data()
}

//------- ATTACK SECTION

type Attack struct{
	Title	string
	Cost	EnergyCost
	Description	string
	BaseDamage	map[string]int//[.]10 -> 10 damage , [+]10 -> +10 damage ...
	Efects	map[string]Efect
}

func (At Attack) SlotData(){
	fmt.Println(At.Title)
}

type Ability struct{
	Title	string
	Cost	EnergyCost
	Description	string
	//BaseDamage	map[string]int//[.]10 -> 10 damage , [+]10 -> +10 damage ...
	Efects	map[string]Efect
}

func (Ab Ability) SlotData(){
	fmt.Println(Ab.Title)
}

type Power struct{
	Title	string
	Cost	EnergyCost
	Description	string
	//BaseDamage	map[string]int//[.]10 -> 10 damage , [+]10 -> +10 damage ...
	Efects	map[string]Efect
}

func (Pw Power) SlotData(){
	fmt.Println(Pw.Title)
}

type CardSlot interface{
	SlotData()
}

//------- EFFECT SECTION

type Coin_Add struct{
	Coins		int
	Operator	string
	Additional	int
	BaseDamage	int
	Objective_Head	Pokemon
	Objective_Tail	Pokemon
}

func (C Coin_Add) Attack(){
	Damage := 0
	for i:=0;i<C.Coins;i++{
		rand.Seed( time.Now().UnixNano())
		Res := rand.Intn(100)%2
		if Res == 0{
			switch(C.Operator){
			case "+":
				Damage = C.BaseDamage + C.Additional
				break
			case "X":
				Damage += C.BaseDamage
			}
		}
	}
	//return Damage
}

type Coin_Stat struct{
	Coins	int
	Head	status
	Tail	status
	Objective_Head	Pokemon
	Objective_Tail	Pokemon
}

func (C Coin_Stat) Attack(){
	rand.Seed( time.Now().UnixNano())
	Res := rand.Intn(100)%2
	if Res == 0{
		//return C.Head
	}else{
		//return C.Tail
	}
}

type Consume_Energy struct{
	Cost	EnergyCost
}

type Counter struct{
	Hp			int
	HpMax		int
	BaseDamage	int
	divisor 	int
	Round		int // 0 = no, 1 = up, -1 = down
}

type Indirect struct{
	Objective	Pokemon
	Damage		int
}

type Flail struct{
	HpMax	int
	Hp		int
	Damage 	int
}

func (S Side)Switch_Pokemon(pos int, Pok Pokemon ) Side{
	Aux := S.Main
	switch(pos){
		case 1:
			S.Main = S.Bench1
			S.Bench1 = Aux
			break
		case 2:
			S.Main = S.Bench2
			S.Bench2 = Aux
			break
		case 3:
			S.Main = S.Bench3
			S.Bench3 = Aux
			break
		case 4:
			S.Main = S.Bench4
			S.Bench4 = Aux
			break
		case 5:
			S.Main = S.Bench5
			S.Bench5 = Aux
			break
	}
	return S
}

type Efect interface{
	Attack()
}

