// Package slices implementa operações eficientes em slices de inteiros.
// Inclui append seguro, reverse, busca linear/binária e sort.
package slices

import (
	"fmt"
	"sort"
)

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

// Sort ordena o slice in-place usando quicksort O(n log n)
func Sort(s []int) {
	sort.Ints(s)
}

// BinarySearch busca binária O(log n) - slice DEVE estar ordenado
func BinarySearch(s []int, target int) int {
	idx := sort.SearchInts(s, target)
	if idx < len(s) && s[idx] == target {
		return idx
	}
	return -1
}

// Max retorna o maior elemento O(n)
func Max(s []int) (int, bool) {
	if len(s) == 0 {
		return 0, false
	}
	max := s[0]
	for _, v := range s[1:] {
		if v > max {
			max = v
		}
	}
	return max, true
}

// Min retorna o menor elemento O(n)
func Min(s []int) (int, bool) {
	if len(s) == 0 {
		return 0, false
	}
	min := s[0]
	for _, v := range s[1:] {
		if v < min {
			min = v
		}
	}
	return min, true
}

// Sum retorna a soma de todos elementos O(n)
func Sum(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}
