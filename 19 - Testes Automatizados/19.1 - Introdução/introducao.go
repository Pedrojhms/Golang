package main

import (
	"fmt"
	"introducao-testes/Enderecos"
)

func main() {
	tipoEndereco := Enderecos.TipoDeEndereco("Avenida Paulista")
	fmt.Println(tipoEndereco)
}
