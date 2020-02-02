package main 

import (
	"log"
	"os"
)
//attack info
type attack struct{
	Name string
	Description string
	Energy map[int]int //[Energy Type] Number of Energies
	Base int //Base damage
	Condition []string
	Coin map[bool]int // [Need Coins?] Numbere of Coins 
	Cost map[int]int // [energy consumed] Number of Energies
	SideEffect string// Poison, Paralisis, Sleep, Frozen, etc...
	
}

//Pokemon main data
type Pokemon struct{
	HpMax 		int			// hp max
	Hp 			int 		// actual hp
	Stage 		int			// Base, stage 1, stage 2 & mega
	Evolve		string		// Evolve to:
	Type 		int 		// Pokemon element
	Slot1 		[3]string	// attack & ability [name, damagebase, cost]
	Slot2 		[3]string 	// attack, ability & empty [name, damagebase, cost]
	Weak 		int			// weakness, damage X2
	Strong 		int			// Strong, damage -30
	Retreat 	[2]int		// retreat cost
	Energy 		[]int 		// energy attched
	Efect1 		int 		// paralisis, sleep & confused
	Efect2 		int 		// burned & poisoned
	Slot 		int			// field position
}

var(
	AttackType = map[string]string{
		"Scratch":"Normal",
		"Dig":"Normal",
		"Mud Slap":"Normal",
		"Pound":"Normal",
		"Solarbeam":"Normal",
		"Bite":"Normal",
		"Tackle":"Normal",
		"Slap":"Normal",
		"Slash":"Normal",
		"Super Psy":"Normal",
		"Smash Kick":"Normal",
		"Flame Tail":"Normal",
		"Gnaw":"Normal",
		"Special Punch":"Normal",
		"Jab":"Normal",
		"Aurora Beam":"Normal",
		"Flare":"Normal",
		"Pot Smash":"Normal",
		"Vine Whip":"Normal",
		"Fire Punch":"Normal",
		"Horn Drill":"Normal",
		"Headbutt":"Normal",
		"Low Kick":"Normal",
		"Rock Throw":"Normal",
	}
)

func (P *Pokemon)Attack(att string, Status *Field, DF *Pokemon, damage int){
	Type,ok := AttackType[att]
	if !ok{
		log.Println("Attack don't found")
		os.Exit(0)
	}
	switch Type {
	case "Normal":
//		Status.Dfn.Hp -= DamageCalc(P,DF,damage)
	}
}
func DamageCalc(At *Pokemon, Df *Pokemon,damage int)int{
	if Df.Strong == At.Type{
		return damage-30
	}else{
		if Df.Weak == At.Type{
			return damage * 2
		}else{
			return damage
		}
	}
}
