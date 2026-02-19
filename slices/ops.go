package slices

// VetorOps: slices como peões go — append xeque, reverse mate
// senai lógico glow-up: o(1) ops nativas + custom

import "fmt"

// AppendSafe adiciona com check cap (evita realloc drama)
func AppendSafe(s []int, val int) []int {
	// append nativo amort o(1), mas log se resize
	if cap(s) == len(s) {
		fmt.Println("resize: cap up como rainha promo")
	}
	return append(s, val)
}

// Reverse custom o(n) — loop manual pra entender internals
func Reverse(s []int) []int {
	rev := make([]int, len(s))
	for i, j := 0, len(s)-1; i <= j; i, j = i+1, j-1 {
		rev[i], rev[j] = s[j], s[i] // swap xadrez
	}
	return rev
}

// IndexOf busca linear o(n) — base pra binária
func IndexOf(s []int, target int) int {
	for i, v := range s {
		if v == target {
			return i // achou!
		}
	}
	return -1 // rei perdido
}

// Copy cria uma cópia shallow do slice
func Copy(s []int) []int {
	if s == nil {
		return nil
	}
	return append([]int(nil), s...)
}

// Subslice retorna uma parte do slice (semelhante a s[start:end])
func Subslice(s []int, start, end int) []int {
	if start < 0 || end > len(s) || start > end {
		return nil
	}
	return s[start:end]
}
