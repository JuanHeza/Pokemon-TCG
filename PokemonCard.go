package main

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
type stage struct {
	stage int // basic = 0, stage 1 = 1, stage 2 = 2
	Prev  int // # of the pokemon in the pokedex
}

func newPokemon() *CardPokemon { return nil }

//PickPokemon :)
func PickPokemon([5]*CardPokemon) *CardPokemon { return nil }

//PickPokemon :)
func (P *Player) PickPokemon() *CardPokemon { return nil }

//PickEnergy :)
func PickEnergy(Ec EnergyCost, cant int, Type Element) Element { return typeColorless }

//PickElement :)
func PickElement() EnergyCost { return EnergyCost{} }

//PickAttack :)
func PickAttack(*CardPokemon) Attacks { return Attacks{} }

//SwitchPokemon :)
func SwitchPokemon(in *CardPokemon, out *CardPokemon) {}

//DiscardEnergy :)
func DiscardEnergy(Ec EnergyCost, cant int, Type Element) EnergyCost { return EnergyCost{} }

/*
laycra company
	fibras artificiales para textiles

presencia global
	asia & lat america

	San Pedro - santa catarina - la fama
	facultad plan empresa?
	gerente area TI

	6 meses
	no posibilidad de estadia (al momento)

	$5200 mes
	caintra
*/
