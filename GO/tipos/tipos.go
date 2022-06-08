package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// numeros inteiros
	fmt.Println(1, 2, 1000)
	// reflect.TypeOf eh o type() do python, retorna o tipo de uma variavel
	fmt.Println("Literal inteiro eh", reflect.TypeOf(32000))
	fmt.Println("O tipo de variavel eh", reflect.TypeOf("TESTE"))

	// sem sinal (so positivos) uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("o byte e", reflect.TypeOf(b))

	// com sinal... int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("O valor maximo do int eh", i1)
	fmt.Println("O tipo de i1 eh", reflect.TypeOf(i1))

	var i2 rune = 'a' // representam um mapeamento da tabela Unicode (int32)
	fmt.Println("O rune eh", reflect.TypeOf(i2))
	fmt.Println(i2)

	// numeros reais (float32, float64)
	var x float32 = 49.99
	fmt.Println("O tipo de x eh", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49.99 eh", reflect.TypeOf(49.99)) // float64

	// boolean
	bo := true
	fmt.Println("O tipo de bo eh", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// string
	s1 := "Ola meu nome eh Henrique"
	fmt.Println(s1 + "!")
	// tamanho da string tipo python
	fmt.Println("O tamanho da string eh", len(s1))

	// string com multiplas linhas
	s2 := `Ola
	meu
	nome
	eh
	Henrique`
	fmt.Println("O tamanho da string eh", len(s2))

	// char????????????
	// nao existe
	// var x char = 'b'
	char := 'a'
	fmt.Println("O valor de char e", char)                // pega o valor
	fmt.Println("O tipo de char e", reflect.TypeOf(char)) // int32
}
