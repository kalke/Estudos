package main

import "fmt"

func main() {
	// Print nao pula linha
	fmt.Print("Mesma")
	fmt.Print("Linha.")
	// Print ln pula linha
	fmt.Println("Nova")
	fmt.Println("Linha.")

	x := 3.141516

	// Nao funciona usando o mais para printar a concatenacao de string com valor numerico
	// fmt.Println(("O valor de X e " + x))
	// Retorna a string de x
	xs := fmt.Sprint(x)
	fmt.Println("O valor de x e " + xs)
	// Para printar o valor de x concatenado sem ter que converte-lo usamos a virgula
	fmt.Println("O valor de x e", x)
	// Printar apenas duas casas decimais do valor de x (arredondamento)
	fmt.Printf("O valor de x e %.2f", x)

	a := 1
	b := 1.999
	c := false
	d := "opa"
	// Os % referenciam o tipo de dado que sera mostrado no console
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	// % v eh uma forma mais generica de se fazer isso
	fmt.Printf("\n%v %v %v %v", a, b, c, d)

}
