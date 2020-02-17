package main

var (
//P is the Attaking Pokemon
//P			*CardPokemon
//Defender is the Defending Pokemon
//Defender 	*CardPokemon
)

//Attacks Struct of the attak in the card slot
type Attacks struct {
	Cost        EnergyCost
	Active      bool
	Name        string
	Description string
	Params      []int
	Ability     bool
	Do          func(*Field, ...int) *Field
}

//DamageCalc is the calculator of weaknesses & resistance damage
func DamageCalc(At *CardPokemon, Df *CardPokemon, damage int) int {
	res, ok := Df.Resistence.Cost[At.Type]
	if ok {
		damage = damage - res
	}
	res, ok = Df.Weaknesses.Cost[At.Type]
	if ok {
		damage = damage * res
	}
	return damage
}

//GetMains eturn the actual main pokemon in both sides
func GetMains(F *Field) (*CardPokemon, *CardPokemon) {
	var P *CardPokemon = F.Player1.Main
	var Defender *CardPokemon = F.Player2.Main
	return P, Defender
}

//Protection Protects a pokemon from damage or status change
type Protection struct {
	Protected           *CardPokemon
	Protection          bool
	DamageProtected     int
	StatDamageProtected bool
	StatMoveProtected   bool
}

//ProtectedPokemon is the ctual Protected pokemon
var ProtectedPokemon Protection

//Shield brings Protection  to the objective pokemon
func Shield(Objective *CardPokemon, Damage int, StatDamage bool, StatMove bool) {
	ProtectedPokemon = Protection{
		Protected:           Objective,
		Protection:          true,
		DamageProtected:     Damage,
		StatDamageProtected: StatDamage,
		StatMoveProtected:   StatMove,
	}
}

// ======================================================================================================================= \\
// ================================================ POKEMON POWER SECTION ================================================ \\
// ======================================================================================================================= \\

//DamageSwap transfer 10 of damage from pokemon P to pokemon Objective
//DamageSwap(F)
func DamageSwap(F *Field, Damages ...int) *Field {
	var P, _ = GetMains(F)
	if P.StatMove == Normal { //sleep, confused, paralyzed?
		Objective := newPokemon()
		for Objective.Hp < Damages[0] {
			Objective = PickPokemon(F.Player2.Bench)
		} //for
		P.Hp += Damages[0]
		if P.Hp > P.MaxHp {
			P.Hp = P.MaxHp
		}
		Objective.Hp -= Damages[0]
	} //if
	return F
}

//EnergyTrans transfer an energy from a pokemon to another
//EnergyTrans(F)
func EnergyTrans(F *Field, Damages ...int) *Field {
	var P, _ = GetMains(F)
	var Donor, Receptor *CardPokemon
	if P.StatMove == Normal {
		Donor = PickPokemon(F.Player1.Bench)
		Receptor = PickPokemon(F.Player1.Bench)
	}
	Donor.Energies.Cost[P.Type]--
	Receptor.Energies.Cost[P.Type]++
	return F
}

//RainDance Provides an extra WaterEnergy to a pokemon
//RainDance(F)
func RainDance(F *Field, Damages ...int) *Field {
	var P, _ = GetMains(F)
	if P.StatMove == Normal {
		Objective := PickPokemon(F.Player2.Bench)
		Objective.Energies.Cost[typeWater]++
	}
	return F
}

//EnergyBurn Temporary changes all energies atched to typeFire
//EnergyBurn(F)
func EnergyBurn(F *Field, Damages ...int) *Field {
	var P, _ = GetMains(F)
	if P.StatMove == Normal {
		//	   P.Energies -> turnEnd
		Count := 0
		for elem, cant := range P.Energies.Cost {
			if elem != typeFire {
				Count += cant
				P.Energies.Cost[elem] -= cant
			}
		}
		P.Energies.Cost[typeFire] += Count
	}
	return F
}

//Buzzap Knock Out Electrode and convert it to a double energy of your choise
//Buzzap(F)
func Buzzap(F *Field, Damages ...int) *Field {
	var P, _ = GetMains(F)
	if P.StatMove == Normal {
		P.Hp = 0
		res := PickPokemon(F.Player1.Bench)
		Qt := PickEnergy(res.Energies, 1, typeColorless)
		res.Energies.Cost[Qt] += 2
		//draw Electrode death
	}
	return F
}

//StrikesBack If Machamp its attacked, the attacker recives 10 of damages
//StrikesBack(F)
func StrikesBack(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	if P.StatMove == Normal {
		Defender.Hp -= Damages[0]
	}
	return F
}

// ================================================================================================================ \\
// ================================================ ATTACK SECTION ================================================ \\
// ================================================================================================================ \\

//ConfuseRay Confuse the defending pokemon
//ConfuseRay(F,#)
func ConfuseRay(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	confused := FlipCoin()
	if confused {
		Defender.StatMove = Confused
	}
	return F
}

//HydroPump does X damage plus Y if if attched 1 energy extra or 2*Y if have 2 or more energies extra
//HydroPump(F,#,#)
func HydroPump(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	Total := Damages[0]
	if P.Energies.Cost[typeWater] > 5 {
		Total += Damages[1]
	} else {
		if P.Energies.Cost[typeWater] > 4 {
			Total += Damages[1]
		}
	}
	Defender.Hp -= DamageCalc(P, Defender, Total)
	return F
}

//Scrunch Creates a damage shield
//Scrunch(F)
func Scrunch(F *Field, Damages ...int) *Field {
	var P, _ = GetMains(F)
	Res := FlipCoin()
	if Res {
		Shield(P, 9999, false, false)
	}
	return F
}

//DoubleEdge damage to both main pokemons
//DoubleEdge(f,#,#)
func DoubleEdge(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	P.Hp -= Damages[1]
	return F
}

//FireSpin discard 2 energies to use this attack
//FireSpin(F,#)
func FireSpin(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	P.Energies = DiscardEnergy(P.Energies, 2, typeColorless)
	Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	return F
}

//Sling sleep the defending pokemon, flip a coin
//Sling(F)
func Sling(F *Field, Damages ...int) *Field {
	var _, Defender = GetMains(F)
	Sleep := FlipCoin()
	if Sleep {
		Defender.StatMove = Asleep
	}
	return F
}

//Metronome copy an attack
//Metronome(F)
func Metronome(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	Red := PickAttack(Defender)
	P.Slot2 = Red
	return F
}

//DragonRage Common Attack
//DragonRage(f,#)
func DragonRage(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	return F
}

//Bubblebeam is a paralyzing attack flip a coin
//Bubblebeam(F,#)
func Bubblebeam(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	paralyzed := FlipCoin()
	if paralyzed {
		Defender.StatMove = Paralyzed
	}
	Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	return F
}

//Jab Common Attack
//Jab(F,#)
func Jab(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	return F
}

//SpecialPunch Common Attack
//SpecialPunch(F,#)
func SpecialPunch(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	return F
}

//SeismicToss Common Attack
//SeismicToss(F,#)
func SeismicToss(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	return F
}

//ThunderWave is a aralyzing attack flip a coin
//ThunderWave(F,#)
func ThunderWave(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	paralyzed := FlipCoin()
	if paralyzed {
		Defender.StatMove = Paralyzed
	}
	Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	return F
}

//SelfDestruct destroys the attacking pokemon and damages both bench and defending pokemon
//SelfDestruct(F,#,#)
func SelfDestruct(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	P.Hp -= 0
	Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	for _, dfb := range F.Player2.Bench {
		dfb.Hp -= Damages[1]
	}
	for _, atb := range F.Player1.Bench {
		atb.Hp -= Damages[1]
	}
	return F
}

//Psychic damages X plus Y*Energies atched to the defending pokemon
//Psychic(F,#,#)
func Psychic(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	Count := 0
	for _, energy := range Defender.Energies.Cost {
		Count += energy
	}
	Total := (Damages[1] * Count) + Damages[0]
	Defender.Hp -= DamageCalc(P, Defender, Total)
	return F
}

//Barrier discard a energy to create a shield of damages and status
//Barrier(F)
func Barrier(F *Field, Damages ...int) *Field {
	var P, _ = GetMains(F)
	P.Energies.Cost[P.Type]--
	Shield(P, 9999, true, true)
	return F
}

//Thrash damages X flip a coin to determine if Y damage the attacking or defending pokemon
//Thrash(F,#,#)
func Thrash(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	Red := FlipCoin()
	Total := Damages[0]
	if Red {
		Total += Damages[1]
	} else {
		P.Hp -= Damages[1]
	}
	Defender.Hp = DamageCalc(P, Defender, Total)
	return F
}

//Toxic Sujper poison, double damage of poison
//Toxic(F,#)
func Toxic(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	Defender.StatDamage = SuperPoison //super poison, damage 20
	return F
}

//Lure switch the defending pokemon to one in the bench
//Lure(F)
func Lure(F *Field, Damages ...int) *Field {
	var _, Defender = GetMains(F)
	var Objective *CardPokemon = nil
	if len(F.Player2.Bench) > 0 {
		Objective = PickPokemon(F.Player2.Bench)
	}
	SwitchPokemon(Defender, Objective)
	return F
}

//FireBlast discard an energy to use this attack
//FireBlast(F,#)
func FireBlast(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	P.Energies.Cost[P.Type]--
	Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	return F
}

//WaterGun damage X plus Y if has 1 extra energy or Y*2 if 2 or more energies
//WaterGun(F)
func WaterGun(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	Total := Damages[0]
	if P.Energies.Cost[P.Type] > 2 {
		Total += Damages[1]
	}
	if P.Energies.Cost[P.Type] > 3 {
		Total += Damages[1]
	}
	Defender.Hp -= DamageCalc(P, Defender, Total)
	return F
}

//Whirlpool the defending pokemon discar an energy
//Whirlpool(F,#)
func Whirlpool(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	Defender.Energies = DiscardEnergy(Defender.Energies, 1, typeColorless)
	return F
}

//Agility damages and create a shield of damages and status
//Agility(F,#)
func Agility(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	Red := FlipCoin()
	if Red {
		Shield(P, 9999, true, true)
	}
	Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	return F
}

//Thunder flip a coin to prevent/recieve a knockback damage Y
//Thunder(F,#,#)
func Thunder(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	Red := FlipCoin()
	if Red {
		P.Hp -= Damages[1]
	}
	Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	return F
}

//Solarbeam Common Attack
//Solarbeam(F,#)
func Solarbeam(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	return F
}

//Thunderbolt discard all energies to use this attack
//Thunderbolt(F,#)
func Thunderbolt(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	P.Energies.Cost = make(map[Element]int)
	Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	return F
}

//Twineedle flip a coin to damages a max of 2 times X
//Twineedle(F,#)
func Twineedle(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	var Res = [2]bool{FlipCoin(), FlipCoin()}
	Total := 0
	for _, At := range Res {
		if At {
			Total += Damages[0]
		}
	}
	Defender.Hp -= DamageCalc(P, Defender, Total)
	return F
}

//PoisonSting Poisoned Attack Flip a foin
//PoisonSting(F,#)
func PoisonSting(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	poison := FlipCoin()
	if poison {
		Defender.StatDamage = Poisoned
	}
	Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	return F
}

//Slam flip a coin to damages a max of 2 times X
//Slam(F,#)
func Slam(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	var Res = [2]bool{FlipCoin(), FlipCoin()}
	Total := 0
	for _, At := range Res {
		if At {
			Total += Damages[0]
		}
	}
	Defender.Hp -= DamageCalc(P, Defender, Total)
	return F
}

//HyperBeam if the defending pokemon has energies, discard one (picked by attacking)
//HyperBeam(F,#)
func HyperBeam(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	if len(Defender.Energies.Cost) > 0 {
		DiscardEnergy(Defender.Energies, 1, typeColorless)
	}
	return F
}

//Slash Common Attack
//Slash(F,#)
func Slash(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	return F
}

//Earthquake damages your benched pokemons
//Earthquake(F,#,#)
func Earthquake(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	for _, atb := range F.Player1.Bench {
		atb.Hp -= Damages[1]
	}
	Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	return F
}

//Thundershock Paralyzing attack flip a coin
//Thundershock(F,#)
func Thundershock(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	paralyzed := FlipCoin()
	if paralyzed {
		Defender.StatMove = Paralyzed
	}
	Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	return F
}

//Thunderpunch damages X flip a coin to determine if Y damage the attacking or defending pokemon
//Thunderpunch(F,#,#)
func Thunderpunch(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	res := FlipCoin()
	Total := Damages[0]
	if res {
		Total += Damages[1]
	} else {
		P.Hp -= Damages[1]
	}
	Defender.Hp -= DamageCalc(P, Defender, Total)
	return F
}

//ElectricShock flip a coin to se if prevent/recive knockback
//ElectricShock(F,#,#)
func ElectricShock(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	res := FlipCoin()
	if res {
		P.Hp -= Damages[1]
	}
	Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	return F
}

//Whirlwind Damages and the adversary changes the pokemon
//Whirlwind(f,#)
func Whirlwind(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	SwitchPokemon(Adversary.PickPokemon(), Defender)
	return F
}

//MirrorMove Copy the damages in the last turn to the defending pokemon
//MirrorMove(F)
func MirrorMove(F *Field, Damages ...int) *Field {
	var _, Defender = GetMains(F)
	Defender.Hp -= LastAttack.Damage
	Defender.StatDamage = LastAttack.StatDamage
	Defender.StatMove = LastAttack.StatMove
	return F
}

//Flamethrower discard an energy to use this attack
//Flamethrower(F,#)
func Flamethrower(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	DiscardEnergy(P.Energies, 1, P.Type)
	Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	return F
}

//TakeDown Damage X and Knockback Y
//TakeDown(F,#,#)
func TakeDown(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	P.Hp -= Damages[1]
	Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	return F
}

//AuroraBeam Common Attack
//AuroraBeam(F,#)
func AuroraBeam(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	return F
}

//IceBeam Paralyzing attack flip a coin
//IceBeam(F,#)
func IceBeam(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	paralyzed := FlipCoin()
	if paralyzed {
		Defender.StatMove = Paralyzed
	}
	Defender.Hp = DamageCalc(P, Defender, Damages[0])
	return F
}

//Pound Common Attack
//Pound(F,#)
func Pound(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	return F
}

//LeekSlap flip a coin, if tails this attack does nothing, anyway the attack is blocked
//LeekSlap(F,#)
func LeekSlap(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	res := FlipCoin()
	if res {
		Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	}
	if F.Player1.Main.Slot1.Name == "Leek Slap" {
		F.Player1.Main.Slot1.Active = false
	} else {
		F.Player1.Main.Slot2.Active = false
	}
	return F
}

//PotSmash Common Attack
//PotSmash(F,#)
func PotSmash(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	return F
}

//Flare Common Attack
//Flare(F,#)
func Flare(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	return F
}

//Hypnosis the defending pokemon is sleeping
//Hypnosis(F)
func Hypnosis(F *Field, Damages ...int) *Field {
	var _, Defender = GetMains(F)
	Defender.StatMove = Asleep
	return F
}

//DreamEater use only if the defending pokemon is sleeping
//DreamEater(F,#)
func DreamEater(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	if Defender.StatMove == Asleep {
		Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	}
	return F
}

//VineWhip Common Attack
//VineWhip(F,#)
func VineWhip(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	return F
}

//Poisonpowder the defending pokemon is Poisoned
//Poisonpowder(F,#)
func Poisonpowder(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	Defender.StatDamage = Poisoned
	Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	return F
}

//DoubleSlap flip 2 coins, the damages = Heads*X
//DoubleSlap(F,#)
func DoubleSlap(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	Total := 0
	res := [2]bool{FlipCoin(), FlipCoin()}
	for _, x := range res {
		if x {
			Total += Damages[0]
		}
	}
	Defender.Hp -= DamageCalc(P, Defender, Total)
	return F
}

//Meditate does damage X plus diference between MaxHp & Hp of the defending pokemon
//Meditate(F,#)
func Meditate(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	Defender.Hp -= DamageCalc(P, Defender, Damages[0]+(Defender.MaxHp-Defender.Hp))
	return F
}

//Recover Discad an energy and get full heal
//Recover(F)
func Recover(F *Field, Damages ...int) *Field {
	var P, _ = GetMains(F)
	P.Energies = DiscardEnergy(P.Energies, 1, P.Type)
	P.Hp = P.MaxHp
	return F
}

//SuperPsy Common Attack
//SuperPsy(F,#)
func SuperPsy(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	return F
}

//Stiffen creates a dmage shield
//Stiffen(F)
func Stiffen(F *Field, Damages ...int) *Field {
	var P, _ = GetMains(F)
	res := FlipCoin()
	if res {
		Shield(P, 9999, false, false)
	}
	return F
}

//KarateChop damage X minus Damagecounter on P * 10
//KarateChop(F,#,#)
func KarateChop(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	Damage := P.MaxHp - P.Hp
	if Damage < Damages[0] {
		Defender.Hp -= DamageCalc(P, Defender, (Damages[0] - Damage))
	}
	return F
}

//Submission Knockback Attack
//Submission(F,#,#)
func Submission(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	P.Hp -= Damages[1]
	return F
}

//Tackle Common Attack
//Tackle(F,#)
func Tackle(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	return F
}

//Flail Damage is equal to damage counters in the pokemon
//Flail(F)
func Flail(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	Defender.Hp -= DamageCalc(P, Defender, P.MaxHp-P.Hp)
	return F
}

//FirePunch Common Attack
//FirePunch(F,#)
func FirePunch(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	return F
}

//DoubleKick flip a coin, damage is equal to Heads * X
//DoubleCick(F,#)
func DoubleKick(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	Total := 0
	res := [2]bool{FlipCoin(), FlipCoin()}
	for _, x := range res {
		if x {
			Total += Damages[0]
		}
	}
	Defender.Hp -= DamageCalc(P, Defender, Total)
	return F
}

//HornDrill Common Attack
//HornDrill(F,#)
func HornDrill(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	return F
}

//Amnesia block an attack of the defending pokemon for 1 turn
//Amnesia(F,#)
//falta agregar por un turno
func Amnesia(F *Field, Damages ...int) *Field {
	var _, Defender = GetMains(F)
	Objective := PickAttack(Defender)
	Objective.Active = false
	return F
}

//Conversion1 change the defender weaknesses to another type
//Conversion1(F)
//falltan cosas
func Conversion1(F *Field, Damages ...int) *Field {
	var _, Defender = GetMains(F)
	if Defender.Weaknesses.Cost != nil {
		Defender.Weaknesses = PickElement()
		//typeColorless NO
	}
	return F
}

//Conversion2 change the attacking pokemon resistance
//Conversion2(F)
//Faltandatos
func Conversion2(F *Field, Damages ...int) *Field {
	var P, _ = GetMains(F)
	P.Resistence = PickElement()
	//typeColorless NO
	return F
}

//Bite Common Attack
//Bite(F,#)
func Bite(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	return F
}

//SuperFang does damage equal to half of the actual defender hp rounded up
//SuperFang(F)
func SuperFang(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	Damage := (Defender.MaxHp - Defender.Hp) / 2
	if Damage%10 != 0 {
		Damage += 5
	}
	Defender.Hp -= DamageCalc(P, Defender, Damage)
	return F
}

//Headbutt Common Attack
//Headbutt(F,#)
func Headbutt(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	return F
}

//Withdraw create a damage shield
//Withdraw(F)
func Withdraw(F *Field, Damages ...int) *Field {
	var P, _ = GetMains(F)
	res := FlipCoin()
	if res {
		Shield(P, 9999, false, false)
	}
	return F
}

//Psyshock Paralyzing Attack flip a coin
//Psyshock(F,#)
func Psyshock(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	paralyzed := FlipCoin()
	if paralyzed {
		Defender.StatMove = Paralyzed
	}
	Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	return F
}

//LeechSeed heal 10hp if the attack is completed
//LeechSeed(F,#)
//faltan datos
func LeechSeed(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	var Damage = 0
	if Damage != 0 {
		P.Hp += 10
	}
	Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	return F
}

//StringShot Paralyzing attack flip a coin
//StringShot(F,#)
func StringShot(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	paralyzed := FlipCoin()
	if paralyzed {
		Defender.StatMove = Paralyzed
	}
	Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	return F
}

//Scratch Common Attack
//Scratch(F,#)
func Scratch(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	return F
}

//Ember discard energy to use this attack
//Ember(F,#)
func Ember(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	P.Energies = DiscardEnergy(P.Energies, 1, P.Type)
	Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	return F
}

//Dig Common Attack
//Dig(F,#)
func Dig(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	return F
}

//MudSlap Common Attack
//MudSlap(F,#)
func MudSlap(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	return F
}

//FuryAttack flip a coin, damages equals to heads * X
//FuryAttack(F,#)
func FuryAttack(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	Total := 0
	res := [2]bool{FlipCoin(), FlipCoin()}
	for _, x := range res {
		if x {
			Total += Damages[0]
		}
	}
	Defender.Hp -= DamageCalc(P, Defender, Total)
	return F
}

//SleepingGas Sleeping Attack flip a coin
//SleepingGas(F,#)
func SleepingGas(F *Field, Damages ...int) *Field {
	var _, Defender = GetMains(F)
	sleep := FlipCoin()
	if sleep {
		Defender.StatMove = Asleep
	}
	return F
}

//DestinyBond Discard an energy to use this attack
//DestinyBond(F)
//faltan datos
func DestinyBond(F *Field, Damages ...int) *Field {
	var P, _ = GetMains(F)
	P.Energies = DiscardEnergy(P.Energies, 1, P.Type)
	// data error
	// if gastly dies in the next turn both pokemon dies
	return F
}

//FoulGas Poisoned and paralyzing attack flip a coin
//FoulGas(F,#)
func FoulGas(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	res := FlipCoin()
	if res {
		Defender.StatDamage = Poisoned
	} else {
		Defender.StatMove = Confused
	}
	Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	return F
}

//LowKick Common Attack
//LowKick(F,#)
func LowKick(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	return F
}

//StunSpore Paralyzing Attack flip a coin
//StunSpore(F,#)
func StunSpore(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	paralyzed := FlipCoin()
	if paralyzed {
		Defender.StatMove = Paralyzed
	}
	Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	return F
}

//HornHazard if tails this attack fail
//HornHazard(F,#)
func HornHazard(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	res := FlipCoin()
	if !res {
		Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	}
	return F
}

//RockThrow Common Attack
//RockThrow(F,#)
func RockThrow(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	return F
}

//Harden creates a damage shield of 30hp
//Harden(F)
func Harden(F *Field, Damages ...int) *Field {
	var P, _ = GetMains(F)
	Shield(P, 30, false, false)
	return F
}

//Gnaw Common Attack
//Gnaw(F,#)
func Gnaw(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	return F
}

//ThunderJolt flip a coin if tail p gets knockback
//ThunderJolt(F,#,#)
func ThunderJolt(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	res := FlipCoin()
	if !res {
		P.Hp -= Damages[1]
	}
	Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	return F
}

//SmashKick Common Attack
//SmashKick(F,#)
func SmashKick(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	return F
}

//FlameTail Common Attack
//FlameTail(F,#)
func FlameTail(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	return F
}

//SandAttack next turn enemy flip a coin, if tail, attack fail
//SandAttack(F,#)
//faltan datos
func SandAttack(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	//Tiene
	//efecto
	//Probabilidad de fallar ataque oponente
	Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	return F
}

//Bubble Paralyzing attack flip a coin
//Bubble(F,#)
func Bubble(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	paralyzed := FlipCoin()
	if paralyzed {
		Defender.StatMove = Paralyzed
	}
	Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	return F
}

//StarFreeze paralyzing attack flip a coin
//StarFreeze(F,#)
func StarFreeze(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	paralyzed := FlipCoin()
	if paralyzed {
		Defender.StatMove = Paralyzed
	}
	Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	return F
}

//Slap Common Attack
//Slap(F,#)
func Slap(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	return F
}

//Bind paralyzing attack
//Bind(F,#)
func Bind(F *Field, Damages ...int) *Field {
	var P, Defender = GetMains(F)
	paralyzed := FlipCoin()
	if paralyzed {
		Defender.StatMove = Paralyzed
	}
	Defender.Hp -= DamageCalc(P, Defender, Damages[0])
	return F
}

/*
func (F *Field, Damages ...int)(*Field) {
	Defender.Hp -= DamageCalc(P,Defender,Damages[0])
	return F
}
*/
