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
	fmt.Println("vetores-go ✨ input: 1 2 3 > ops")

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("vetor> ")
	scanner.Scan()
	input := strings.Fields(scanner.Text())

	s := make([]int, len(input))
	for i, str := range input {
		s[i], _ = strconv.Atoi(str) // error handle pro
	}

	fmt.Printf("len/cap: %d/%d\n", len(s), cap(s))
	s = slices.AppendSafe(s, 99)
	fmt.Println("append 99:", s)
	fmt.Println("reverse:", slices.Reverse(s))
	fmt.Println("index 2:", slices.IndexOf(s, 2))
}
