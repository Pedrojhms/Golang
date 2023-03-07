package main

import (
	"Golang/auxiliar"
	"fmt"
	"github.com/badoux/checkmail"
)

func main() {
	auxiliar.Escrever()
	fmt.Println("Escrevendo do arquivo main")

	err := checkmail.ValidateFormat("devbook@gmail.com")
	fmt.Println(err)
}
