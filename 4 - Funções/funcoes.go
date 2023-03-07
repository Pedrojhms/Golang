package main

import "fmt"

func somar(n1 int, n2 int) int {
	return n1 + n2
}

func subtrair(n1 int, n2 int) int {
	return n1 - n2
}

func calculosMatematicos(n1, n2 int) (int, int) {
	return somar(n1, n2), subtrair(n1, n2)
}

func main() {
	soma := somar(1, 5)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt + "!"
	}

	resultado := f("Função f")
	fmt.Println(resultado)

	resultadoSoma, resultadoSubtracao := calculosMatematicos(5, 5)
	fmt.Println(resultadoSoma, resultadoSubtracao)
}
