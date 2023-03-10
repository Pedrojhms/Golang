package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json: "nome"`
	Raca  string `json :"raca"`
	Idade uint   `json: "idade"`
}

func main() {
	c := cachorro{"Rex", "Dálmata", 3}
	fmt.Println(c)

	cachorroEmJson, err := json.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cachorroEmJson)
	fmt.Println(bytes.NewBuffer(cachorroEmJson))

	c2 := map[string]string{
		"nome": "Toby",
		"raca": "Poodle",
	}

	cachorro2EmJson, err := json.Marshal(c2)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cachorro2EmJson)
	fmt.Println(bytes.NewBuffer(cachorro2EmJson))
}
