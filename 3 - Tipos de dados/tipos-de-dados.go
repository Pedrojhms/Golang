package main

import (
	"errors"
	"fmt"
)

func main() {
	numero1 := -10000000000000
	fmt.Println(numero1)

	var numero2 uint32 = 10000
	fmt.Println(numero2)

	//alias
	//INT32 = RUNE
	var numero3 rune = 123456
	fmt.Println(numero3)

	//BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)

	char := 'A'
	fmt.Println(char)

	texto := 5
	fmt.Println(texto)

	var booleano1 bool
	fmt.Println(booleano1)

	var err error = errors.New("Erro interno")
	fmt.Println(err)
}
