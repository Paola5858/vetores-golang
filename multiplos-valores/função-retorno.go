package main

import "fmt"

func analisarNotas(nota1, nota2 float64) (float64, string) {
	media := (nota1 + nota2) / 2
	if media >= 6 {
		return media, "Aprovado"
	} else {
		return media, "Reprovado"
		}
} 
func main() {
	nota1 := 7.5 
	nota2 := 5.5 
	media, resultado := analisarNotas(nota1, nota2)
	fmt.Printf("A média é: %.2f e o resultado é: %s\n", media, resultado)
}
