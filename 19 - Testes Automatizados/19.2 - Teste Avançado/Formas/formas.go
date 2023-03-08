package formas

import (
	"fmt"
	"math"
)

type Forma interface {
	area() float64
}

type Retangulo struct {
	Altura  float64
	Largura float64
}

type Circulo struct {
	raio float64
}

func (r Retangulo) Area() float64 {
	return r.Altura * r.Largura
}

func (c Circulo) Area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func main() {
	r := Retangulo{10, 15}

	c := Circulo{10}

	fmt.Println(r, c)
}
