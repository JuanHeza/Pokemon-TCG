package main

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
)

var (
	v       []map[string]interface{}
	card    = []CardBase{}
	pokemon = []CardPokemon{}
	trainer = []CardTrainer{}
	energy  = []CardEnergy{}
)

//Decode the data from a csv
func Decode(file string) {
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
		switch carta["supertype"] {
		case "Energy":
			log.Println("Energy") //, carta)
			break
		case "Trainer":
			log.Println("Trainer") //, carta)
			break
		case "Pokémon":
			log.Println("Pokemon") //, carta)
			newPokemonCard(carta)
			break
		default:
			log.Println("Error", carta)
			break
		}
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
