package Enderecos_test

import (
	. "introducao-testes/Enderecos"
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	t.Parallel()

	cenarioDeTeste := []cenarioDeTeste{
		{"Rua ABC", "RUA"},
		{"Avenida Paulista", "AVENIDA"},
		{"Rodovia dos Imigrantes", "RODOVIA"},
		{"Praça das Rosas", "Tipo Inválido"},
		{"Estrada Qualquer", "ESTRADA"},
		{"RUA DOS BOBOS", "RUA"},
		{"AVENIDA REBOUÇAS", "AVENIDA"},
		{"", "Tipo Inválido"},
	}

	for _, cenario := range cenarioDeTeste {
		retornoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if retornoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido é diferente do esperado! Esperado %s, recebido %s", cenario.retornoEsperado, retornoRecebido)
		}
	}
}

func TestQualquer(t *testing.T) {
	t.Parallel()

	if 1 > 2 {
		t.Errorf("Teste Quebrou!")
	}
}
