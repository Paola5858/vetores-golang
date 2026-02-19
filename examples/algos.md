# Algoritmos de Busca em Go

## Busca Linear O(n)

A busca linear percorre cada elemento até encontrar o alvo.

```
go
func LinearSearch(s []int, target int) int {
    for i, v := range s {
        if v == target {
            return i
        }
    }
    return -1
}
```

**Complexidade:** O(n) - pior caso percorre todos os elementos

## Busca Binária O(log n)

A busca binária requer um array ordenado e divide o espaço de busca pela metade a cada iteração.

```
go
func BinarySearch(s []int, target int) int {
    left, right := 0, len(s)-1
    
    for left <= right {
        mid := left + (right-left)/2 // evita overflow
        
        if s[mid] == target {
            return mid
        } else if s[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return -1
}
```

**Complexidade:** O(log n) - a cada iteração reduz pela metade

**Pré-requisitos:** Array ordenado

## Exemplo de Uso

```
go
sorted := []int{1, 3, 5, 7, 9, 11, 13, 15}

// Busca linear
idx := LinearSearch(sorted, 7) // retorna 3

// Busca binária (mais rápida para arrays grandes)
idx := BinarySearch(sorted, 7) // retorna 3
```

## Quando Usar Cada Um

| Cenário | Algoritmo |
|---------|-----------|
| Array pequeno (< 50 elementos) | Linear |
| Array ordenado grande | Binária |
| Busca frequente | Binária |
| Array não ordenado | Linear |
