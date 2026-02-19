# Exemplos de Uso - vetores-golang

## 🎯 Operações Básicas

### Append Seguro
```go
s := []int{1, 2, 3}
s = slices.AppendSafe(s, 4)
// Output: [1, 2, 3, 4]
// Se cap estiver cheio: "resize: cap up como rainha promo"
```

### Reverse
```go
s := []int{1, 2, 3, 4, 5}
rev := slices.Reverse(s)
// Output: [5, 4, 3, 2, 1]
```

### Busca Linear
```go
s := []int{10, 20, 30, 40}
idx := slices.IndexOf(s, 30)
// Output: 2
```

## 🚀 Operações Avançadas

### Sort + Binary Search
```go
s := []int{5, 2, 8, 1, 9}
slices.Sort(s)
// s agora: [1, 2, 5, 8, 9]

idx := slices.BinarySearch(s, 8)
// Output: 3 (O(log n) - muito mais rápido!)
```

### Estatísticas
```go
s := []int{15, 3, 42, 7, 23}

max, _ := slices.Max(s)    // 42
min, _ := slices.Min(s)    // 3
sum := slices.Sum(s)       // 90
```

### Copy & Subslice
```go
original := []int{1, 2, 3, 4, 5}

// Cópia independente
copia := slices.Copy(original)
copia[0] = 999
// original ainda é [1, 2, 3, 4, 5]

// Subslice
sub := slices.Subslice(original, 1, 4)
// Output: [2, 3, 4]
```

## 💡 Caso de Uso Real: Análise de Notas

```go
package main

import (
    "fmt"
    "github.com/Paola5858/vetores-golang/slices"
)

func main() {
    notas := []int{85, 92, 78, 95, 88, 76, 90}
    
    // Ordenar
    slices.Sort(notas)
    fmt.Println("Notas ordenadas:", notas)
    
    // Estatísticas
    max, _ := slices.Max(notas)
    min, _ := slices.Min(notas)
    sum := slices.Sum(notas)
    media := float64(sum) / float64(len(notas))
    
    fmt.Printf("Maior nota: %d\n", max)
    fmt.Printf("Menor nota: %d\n", min)
    fmt.Printf("Média: %.2f\n", media)
    
    // Buscar nota específica
    if idx := slices.BinarySearch(notas, 90); idx != -1 {
        fmt.Printf("Nota 90 encontrada na posição %d\n", idx)
    }
}
```

## ⚡ Complexidade Big-O

| Operação | Complexidade | Quando Usar |
|----------|--------------|-------------|
| AppendSafe | O(1) amortizado | Adicionar elementos |
| Reverse | O(n) | Inverter ordem |
| IndexOf | O(n) | Busca em slice não ordenado |
| Sort | O(n log n) | Ordenar dados |
| BinarySearch | O(log n) | Busca em slice ordenado |
| Max/Min | O(n) | Encontrar extremos |
| Sum | O(n) | Somar elementos |

## 🎓 Dicas Pro

1. **Use BinarySearch apenas em slices ordenados**
   ```go
   s := []int{5, 2, 8}
   slices.Sort(s)  // SEMPRE ordene antes!
   idx := slices.BinarySearch(s, 5)
   ```

2. **Copy antes de modificar**
   ```go
   original := []int{1, 2, 3}
   backup := slices.Copy(original)
   slices.Sort(original)  // backup não é afetado
   ```

3. **Verifique retornos de Max/Min**
   ```go
   if max, ok := slices.Max(s); ok {
       fmt.Println("Max:", max)
   } else {
       fmt.Println("Slice vazio!")
   }
   ```

---
✨ **paola** 💋 | senai lógico → algos reais
