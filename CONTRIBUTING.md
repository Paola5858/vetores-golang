# 🤝 Guia de Contribuição

Obrigada por considerar contribuir com o **vetores-golang**! Este documento fornece diretrizes para tornar o processo de contribuição claro e eficiente.

## 📋 Código de Conduta

- Seja respeitoso e inclusivo
- Aceite críticas construtivas
- Foque no que é melhor para a comunidade
- Mostre empatia com outros colaboradores

## 🚀 Como Contribuir

### 1. Reportar Bugs

Encontrou um bug? Abra uma issue com:
- **Título descritivo**
- **Passos para reproduzir**
- **Comportamento esperado vs atual**
- **Versão do Go** (`go version`)
- **Sistema operacional**

Exemplo:
```
Título: BinarySearch retorna índice errado para slice com duplicatas

Passos:
1. s := []int{1, 2, 2, 3}
2. idx := slices.BinarySearch(s, 2)
3. Retorna 1, mas deveria retornar 2

Go: 1.22
OS: Windows 11
```

### 2. Sugerir Funcionalidades

Tem uma ideia? Abra uma issue com:
- **Descrição clara** da funcionalidade
- **Caso de uso** prático
- **Exemplo de código** (se possível)

### 3. Contribuir com Código

#### Setup do ambiente
```bash
# Fork o repositório no GitHub
git clone https://github.com/SEU_USUARIO/vetores-golang.git
cd vetores-golang
go mod tidy
```

#### Criar branch
```bash
git checkout -b feature/nome-da-funcionalidade
# ou
git checkout -b fix/nome-do-bug
```

#### Fazer alterações
1. **Escreva código limpo**
   - Use `go fmt` para formatar
   - Siga convenções Go (Effective Go)
   - Nomes descritivos em inglês

2. **Adicione testes**
   ```go
   func TestNovaFuncao(t *testing.T) {
       tests := []struct {
           name string
           input []int
           want []int
       }{
           {"caso básico", []int{1, 2}, []int{2, 1}},
           {"vazio", []int{}, []int{}},
       }
       
       for _, tt := range tests {
           t.Run(tt.name, func(t *testing.T) {
               got := NovaFuncao(tt.input)
               assert.Equal(t, tt.want, got)
           })
       }
   }
   ```

3. **Documente funções públicas**
   ```go
   // NovaFuncao faz X com o slice s.
   // Retorna Y ou erro se Z.
   //
   // Exemplo:
   //   s := []int{1, 2, 3}
   //   result := NovaFuncao(s)
   //   // result: [3, 2, 1]
   func NovaFuncao(s []int) []int {
       // implementação
   }
   ```

#### Verificar qualidade
```bash
# Formatar código
go fmt ./...

# Verificar erros comuns
go vet ./...

# Rodar testes
go test ./slices -v

# Verificar cobertura (manter > 90%)
go test ./slices -cover
```

#### Commit
Use commits semânticos:
```bash
git commit -m "feat: adiciona função Filter para slices"
git commit -m "fix: corrige BinarySearch com duplicatas"
git commit -m "docs: atualiza README com exemplos"
git commit -m "test: adiciona casos de borda para Sort"
git commit -m "refactor: simplifica lógica de Reverse"
```

Prefixos:
- `feat`: nova funcionalidade
- `fix`: correção de bug
- `docs`: documentação
- `test`: testes
- `refactor`: refatoração sem mudar comportamento
- `perf`: melhoria de performance
- `chore`: tarefas de manutenção

#### Push e Pull Request
```bash
git push origin feature/nome-da-funcionalidade
```

No GitHub:
1. Abra Pull Request
2. Descreva as mudanças
3. Referencie issues relacionadas (#123)
4. Aguarde review

## ✅ Checklist de PR

Antes de submeter, verifique:

- [ ] Código formatado (`go fmt`)
- [ ] Sem erros (`go vet`)
- [ ] Testes passando (`go test ./...`)
- [ ] Cobertura > 90% (`go test -cover`)
- [ ] Funções públicas documentadas
- [ ] Commits semânticos
- [ ] README atualizado (se necessário)
- [ ] Exemplos adicionados (se nova feature)

## 🎯 Áreas que Precisam de Ajuda

### Alta prioridade
- [ ] Implementar generics (Go 1.18+)
- [ ] Adicionar benchmarks
- [ ] Funções: Filter, Map, Reduce

### Média prioridade
- [ ] Melhorar documentação
- [ ] Adicionar mais exemplos
- [ ] Otimizar performance

### Baixa prioridade
- [ ] Suporte a outros tipos (float64, string)
- [ ] CLI interativa
- [ ] Visualização de algoritmos

## 📚 Recursos

- [Effective Go](https://go.dev/doc/effective_go)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [Table Driven Tests](https://dave.cheney.net/2019/05/07/prefer-table-driven-tests)

## 💬 Dúvidas?

- Abra uma issue com a tag `question`
- Entre em contato: [@Paola5858](https://github.com/Paola5858)

---

Obrigada por contribuir! 🎉
