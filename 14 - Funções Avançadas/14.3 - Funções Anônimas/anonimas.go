package main

import "fmt"

func main() {

	retorno := func(texto string, numero int) string {
		return fmt.Sprintf("Recebido -> %s %d", texto, numero)
	}("Passando Parâmetro", 0)

	fmt.Println(retorno)
}
