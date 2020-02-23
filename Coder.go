package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var (
	v       []map[string]interface{}
	card    = []CardBase{}
	pokemon = []CardPokemon{}
	trainer = []CardTrainer{}
	energy  = []CardEnergy{}
	//Sets store the info of each set the first index is the set and the second the ID of the card
	Sets = map[string](map[string]interface{}){}
)

//Decode the data from a csv
func Decode(file string) {
	var card interface{}

	log.Println("Decode")
	dat, err := os.Open(file)
	if err != nil {
		log.Println(err)
		return
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(dat)
	//dec := json.NewDecoder(dat)
	err = json.Unmarshal(buf.Bytes(), &v)
	if err != nil {
		log.Println(err)
		return
	}
	/*if err := dec.Decode(&v); err != nil {
		log.Println(err)
		return
	}
	*/
	for _, carta := range v {
		_, ok := Sets[fmt.Sprintf("%s", carta["set"])]
		if !ok {
			Sets[fmt.Sprintf("%s", carta["set"])] = make(map[string]interface{})
		}
		switch carta["supertype"] {
		case "Energy":
			//log.Println("Energy") //, carta)
			card = newEnergyCard(carta)
			Sets[card.(CardEnergy).Set][card.(CardEnergy).ID] = card
			break
		case "Trainer":
			//log.Println("Trainer") //, carta)
			card = newTrainerCard(carta)
			Sets[card.(CardTrainer).Set][card.(CardTrainer).ID] = card
			break
		case "Pokémon":
			//log.Println("Pokemon") //, carta)
			card = newPokemonCard(carta)
			Sets[card.(CardPokemon).Set][card.(CardPokemon).ID] = card
			break
		default:
			log.Println("Error", carta)
			break
		}
		//log.Println(card)
	}
}

//Encode the data to a csv
func Encode(file string) {
	log.Println("Encode")
	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(&v); err != nil {
		log.Println(err)
	}
}

/*
    for {
        var v map[string]interface{}
        if err := dec.Decode(&v); err != nil {
            log.Println(err)
            return
        }
        for k := range v {
            if k != "Name" {
                delete(v, k)
            }
        }

	}
*/
