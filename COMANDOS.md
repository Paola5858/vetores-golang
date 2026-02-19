# 🛠️ Comandos Úteis - vetores-golang

## 🚀 Desenvolvimento

### Rodar o programa
```bash
go run main.go
# Input: 5 2 8 1 9
```

### Rodar com input direto
```bash
echo 5 2 8 1 9 | go run main.go
```

## 🧪 Testes

### Rodar todos os testes
```bash
go test ./...
```

### Rodar testes com verbose
```bash
go test ./slices -v
```

### Rodar testes com cobertura
```bash
go test ./slices -cover
```

### Gerar relatório de cobertura
```bash
go test ./slices -coverprofile=coverage.out
go tool cover -html=coverage.out
```

### Rodar teste específico
```bash
go test ./slices -run TestBinarySearch
```

## 📦 Dependências

### Instalar/atualizar dependências
```bash
go mod tidy
```

### Verificar dependências
```bash
go mod verify
```

### Listar dependências
```bash
go list -m all
```

## 🔍 Análise de Código

### Formatar código
```bash
go fmt ./...
```

### Verificar erros comuns
```bash
go vet ./...
```

### Linter (instalar golangci-lint antes)
```bash
golangci-lint run
```

## 📊 Benchmarks

### Criar benchmark (exemplo)
```go
// slices/ops_bench_test.go
func BenchmarkSort(b *testing.B) {
    s := []int{5, 2, 8, 1, 9, 3, 7, 4, 6}
    for i := 0; i < b.N; i++ {
        Sort(s)
    }
}
```

### Rodar benchmarks
```bash
go test ./slices -bench=.
```

### Benchmark com memória
```bash
go test ./slices -bench=. -benchmem
```

## 🏗️ Build

### Build executável
```bash
go build -o vetores.exe
```

### Build otimizado (produção)
```bash
go build -ldflags="-s -w" -o vetores.exe
```

### Build para Linux (cross-compile)
```bash
GOOS=linux GOARCH=amd64 go build -o vetores
```

## 📝 Documentação

### Gerar documentação local
```bash
go doc -all ./slices
```

### Ver doc de função específica
```bash
go doc slices.BinarySearch
```

### Servidor de documentação
```bash
godoc -http=:6060
# Acesse: http://localhost:6060/pkg/github.com/Paola5858/vetores-golang/slices/
```

## 🐛 Debug

### Rodar com race detector
```bash
go test ./slices -race
```

### Profile CPU
```bash
go test ./slices -cpuprofile=cpu.prof
go tool pprof cpu.prof
```

### Profile memória
```bash
go test ./slices -memprofile=mem.prof
go tool pprof mem.prof
```

## 🔄 Git

### Commit com mensagem descritiva
```bash
git add .
git commit -m "feat: adiciona Sort e BinarySearch com 92.5% coverage"
```

### Push e trigger CI
```bash
git push origin main
# CI roda automaticamente no GitHub Actions
```

### Ver status do CI
```bash
# Acesse: https://github.com/Paola5858/vetores-golang/actions
```

## 📈 Métricas

### Contar linhas de código
```bash
# Windows
dir /s /b *.go | findstr /v "_test.go" | xargs wc -l

# Linux/Mac
find . -name "*.go" ! -name "*_test.go" | xargs wc -l
```

### Complexidade ciclomática (instalar gocyclo)
```bash
gocyclo -over 10 .
```

## 🎯 Atalhos Úteis

### Limpar cache de testes
```bash
go clean -testcache
```

### Limpar build cache
```bash
go clean -cache
```

### Atualizar Go modules
```bash
go get -u ./...
go mod tidy
```

---
✨ **paola** 💋 | comandos pro workflow eficiente
