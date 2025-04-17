package main

import "fmt"

func main() {
    var numeros [5]int = [5]int{10, 20, 30, 40, 50}
    fmt.Println("Elementos do vetor:")
    for _, valor := range numeros {
        fmt.Println(valor)
    }
}
// O código acima cria um vetor de inteiros com 5 elementos e imprime cada elemento do vetor usando um loop for. O uso de range permite iterar sobre os elementos do vetor sem precisar usar índices. Isso torna o código mais legível e menos propenso a erros. Além disso, o uso de fmt.Println permite imprimir os valores de forma formatada, facilitando a leitura dos resultados.