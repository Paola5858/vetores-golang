package main

import "fmt"

func main() {
    var frutas []string
    frutas = append(frutas, "Maçã", "Banana", "Laranja")
    
    fmt.Println("Frutas no slice:")
    for _, fruta := range frutas {
        fmt.Println(fruta)
    }
}
// O código acima cria um slice de strings chamado "frutas" e adiciona três frutas a ele. Em seguida, imprime cada fruta no console.