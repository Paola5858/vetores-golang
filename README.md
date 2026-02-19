# vetores-golang ✨ go slices como xeque-mate

Go vetores/slices: append, index, reverse — eficiente O(1) ops. senai lógico pro, mas scale pra algos reais (busca, sort).

[![go](https://img.shields.io/badge/go-1.22-green)](go.mod) [![test](https://github.com/Paola5858/vetores-golang/workflows/CI/badge.svg)](actions)

## por quê brilha 🔥
| op | big-o | exemplo |
|----|-------|---------|
| append | amort o(1) | `s = append(s, 42)` |
| len/cap | o(1) | `fmt.Println(len(s))` |
| slice sub | o(1) | `sub := s[1:3]` |
| reverse | o(n) | custom fn |

## roda em 10s (ubuntu/win/mac)
```bash
git clone url && cd vetores-golang
go mod tidy
go run main.go  # interativo: input [1 2 3] > ops
```

## códigos refatorados: go clean (ryanmcdermott inspirado)

### slices/ops.go
Funções puras, comentários story:

```
go
package slices

// AppendSafe adiciona com check cap (evita realloc drama)
func AppendSafe(s []int, val int) []int {
    if cap(s) == len(s) {
        fmt.Println("resize: cap up como rainha promo")
    }
    return append(s, val)
}

// Reverse custom o(n) — loop manual pra entender internals
func Reverse(s []int) []int {
    rev := make([]int, len(s))
    for i, j := 0, len(s)-1; i <= j; i, j = i+1, j-1 {
        rev[i], rev[j] = s[j], s[i]
    }
    return rev
}

// IndexOf busca linear o(n) — base pra binária
func IndexOf(s []int, target int) int {
    for i, v := range s {
        if v == target {
            return i
        }
    }
    return -1
}
```

### tests/ops_test.go
Table-driven como go best practices:

```
go
func TestAppendSafe(t *testing.T) {
    tests := []struct {
        name string
        s    []int
        val  int
        want []int
    }{
        {"basic", []int{1, 2}, 3, []int{1, 2, 3}},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := slices.AppendSafe(tt.s, tt.val)
            assert.Equal(t, tt.want, got)
        })
    }
}
```

## roadmap (kamranahmedse vibe)

- v1.1: sort + busca binária
- v2: concorrência channels

// glitter + goroutines paola 💋🎀
