package main

import "fmt"

func main() {
    EstoqueCantina := map[string]float64{
        "Refrigerante":  10,
        "Pão de queijo": 15,
        "Coxinha":       20,
    }

    EstoqueCantina["Coxinha"] -= 2

    EstoqueCantina["Pão de queijo"] -= 1

    for k, v := range EstoqueCantina {
        fmt.Printf("Oii diva, então esse foi o produto que você pegou: %s, E essa, a quantidade que sobrou pro seus colegas: %.2f\n", k, v)
    }
}