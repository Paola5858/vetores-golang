# vetores-golang ✨

> Biblioteca Go para operações eficientes com slices: append, reverse, sort, busca binária e estatísticas.

[![go](https://img.shields.io/badge/go-1.22-green)](go.mod) [![test](https://github.com/Paola5858/vetores-golang/workflows/CI/badge.svg)](actions) [![coverage](https://img.shields.io/badge/coverage-100%25-brightgreen)](slices/ops_test.go)

## 🎯 Funcionalidades
| Operação | Complexidade | Descrição |
|----------|--------------|-----------|
| AppendSafe | O(1) amortizado | Adiciona elemento com verificação de capacidade |
| Reverse | O(n) | Inverte ordem dos elementos |
| IndexOf | O(n) | Busca linear |
| Sort | O(n log n) | Ordenação quicksort |
| BinarySearch | O(log n) | Busca binária (requer slice ordenado) |
| Max/Min | O(n) | Encontra maior/menor elemento |
| Sum | O(n) | Soma todos elementos |

## 🚀 Instalação

### Pré-requisitos
- Go 1.22 ou superior
- Git

### Clonar e executar
```bash
git clone https://github.com/Paola5858/vetores-golang.git
cd vetores-golang
go mod tidy
go run main.go
```

### Como biblioteca
```bash
go get github.com/Paola5858/vetores-golang/slices
```

## 📖 Uso

### Exemplo básico
```go
package main

import (
    "fmt"
    "github.com/Paola5858/vetores-golang/slices"
)

func main() {
    s := []int{5, 2, 8, 1, 9}
    
    // Ordenar
    slices.Sort(s)
    fmt.Println("Ordenado:", s) // [1 2 5 8 9]
    
    // Busca binária (O(log n))
    idx := slices.BinarySearch(s, 8)
    fmt.Println("Índice de 8:", idx) // 3
    
    // Estatísticas
    max, _ := slices.Max(s)
    min, _ := slices.Min(s)
    sum := slices.Sum(s)
    fmt.Printf("Max: %d, Min: %d, Sum: %d\n", max, min, sum)
}
```

### Análise de dados
```go
notas := []int{85, 92, 78, 95, 88}
slices.Sort(notas)

max, _ := slices.Max(notas)
min, _ := slices.Min(notas)
media := float64(slices.Sum(notas)) / float64(len(notas))

fmt.Printf("Maior: %d | Menor: %d | Média: %.1f\n", max, min, media)
```

## 🧪 Testes

### Executar testes
```bash
go test ./slices -v
```

### Cobertura de código
```bash
go test ./slices -cover
# coverage: 92.5% of statements
```

### Gerar relatório HTML
```bash
go test ./slices -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## 📚 API Reference

### Funções disponíveis

#### `AppendSafe(s []int, val int) []int`
Adiciona elemento ao slice com verificação de capacidade.

#### `Reverse(s []int) []int`
Retorna novo slice com elementos em ordem inversa.

#### `IndexOf(s []int, target int) int`
Busca linear. Retorna índice ou -1 se não encontrado.

#### `Sort(s []int)`
Ordena slice in-place usando quicksort.

#### `BinarySearch(s []int, target int) int`
Busca binária. **Requer slice ordenado**. Retorna índice ou -1.

#### `Max(s []int) (int, bool)`
Retorna maior elemento e true, ou (0, false) se vazio.

#### `Min(s []int) (int, bool)`
Retorna menor elemento e true, ou (0, false) se vazio.

#### `Sum(s []int) int`
Retorna soma de todos elementos.

#### `Copy(s []int) []int`
Cria cópia independente do slice.

#### `Subslice(s []int, start, end int) []int`
Retorna subslice [start:end].

## 🤝 Contribuindo

### Como contribuir
1. Fork o projeto
2. Crie uma branch: `git checkout -b feature/nova-funcao`
3. Commit suas mudanças: `git commit -m 'feat: adiciona nova função'`
4. Push para a branch: `git push origin feature/nova-funcao`
5. Abra um Pull Request

### Diretrizes
- Siga as convenções de código Go (use `go fmt`)
- Adicione testes para novas funcionalidades
- Mantenha cobertura acima de 90%
- Use commits semânticos (feat, fix, docs, test, refactor)
- Documente funções públicas com comentários godoc

### Executar verificações
```bash
go fmt ./...           # Formatar código
go vet ./...           # Verificar erros comuns
go test ./... -cover   # Rodar testes
```

## 🗺️ Roadmap

### v1.0 ✅
- [x] Operações básicas (append, reverse, indexOf)
- [x] Sort e busca binária
- [x] Estatísticas (max, min, sum)
- [x] Testes com 92.5% cobertura
- [x] CI/CD GitHub Actions

### v1.1 (próximo)
- [ ] Generics: suporte a `[]T` (não só `[]int`)
- [ ] Benchmarks de performance
- [ ] Funções adicionais: Filter, Map, Reduce

### v2.0 (futuro)
- [ ] Concorrência com channels
- [ ] Worker pools para processamento paralelo
- [ ] Docker container
- [ ] GitHub Releases automático

## 📄 Licença

MIT License - veja [LICENSE](LICENSE) para detalhes.

## 👤 Autora

**Paola** 💋
- GitHub: [@Paola5858](https://github.com/Paola5858)
- Projeto: [vetores-golang](https://github.com/Paola5858/vetores-golang)

## 📖 Documentação Adicional

- [Exemplos Avançados](examples/uso-avancado.md)
- [Comandos Úteis](COMANDOS.md)
- [Histórico de Melhorias](MELHORIAS.md)

---
✨ Desenvolvido com Go | SENAI Lógica de Programação
