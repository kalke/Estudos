package main

import (
	"fmt"
	"math"
)

func main() {

	// apos a definicao da varivael sou obrigado a usa-la
	const PI float64 = 3.1415
	var raio = 3.2 // tipo (float64) inferido pelo compilador

	// forma reduzida de criar uma variavel
	area := PI * math.Pow(raio, 2)
	fmt.Println("A area da circunferencia e:", area)

	// bloco de construcao
	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)
	fmt.Println(a, b, c, d)

	// varias variaveis na mesma linha
	var e, f bool = true, false
	fmt.Println(e, f)

	g, h, i := 2, false, "epa!"
	fmt.Println(g, h, i)

}
