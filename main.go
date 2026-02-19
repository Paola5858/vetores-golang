package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Paola5858/vetores-golang/slices"
)

func main() {
	fmt.Println("✨ vetores-go | input: 1 2 3 > ops")

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("vetor> ")
	if !scanner.Scan() {
		fmt.Println("❌ erro lendo input")
		return
	}

	input := strings.Fields(scanner.Text())
	if len(input) == 0 {
		fmt.Println("❌ vetor vazio")
		return
	}

	s := make([]int, 0, len(input))
	for _, str := range input {
		val, err := strconv.Atoi(str)
		if err != nil {
			fmt.Printf("⚠️ ignorando '%s' (não é número)\n", str)
			continue
		}
		s = append(s, val)
	}

	fmt.Printf("\n📊 len/cap: %d/%d\n", len(s), cap(s))
	
	// Append
	s = slices.AppendSafe(s, 99)
	fmt.Println("➕ append 99:", s)
	
	// Reverse
	fmt.Println("🔄 reverse:", slices.Reverse(s))
	
	// IndexOf
	if idx := slices.IndexOf(s, 2); idx != -1 {
		fmt.Printf("🔍 index de 2: %d\n", idx)
	} else {
		fmt.Println("🔍 2 não encontrado")
	}
	
	// Sort + BinarySearch
	slices.Sort(s)
	fmt.Println("🔢 sorted:", s)
	if idx := slices.BinarySearch(s, 99); idx != -1 {
		fmt.Printf("🎯 busca binária 99: index %d\n", idx)
	}
	
	// Stats
	if max, ok := slices.Max(s); ok {
		fmt.Printf("🔺 max: %d\n", max)
	}
	if min, ok := slices.Min(s); ok {
		fmt.Printf("🔻 min: %d\n", min)
	}
	fmt.Printf("➕ sum: %d\n", slices.Sum(s))
	
	fmt.Println("\n✨ done! paola 💋")
}
