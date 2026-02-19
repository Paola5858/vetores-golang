package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppendSafe(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		val  int
		want []int
	}{
		{"basic", []int{1, 2}, 3, []int{1, 2, 3}},
		{"full cap", []int{1}, 2, []int{1, 2}},
		{"empty", []int{}, 1, []int{1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AppendSafe(tt.s, tt.val)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		want []int
	}{
		{"odd length", []int{1, 2, 3}, []int{3, 2, 1}},
		{"even length", []int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
		{"single element", []int{1}, []int{1}},
		{"empty", []int{}, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Reverse(tt.s)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestIndexOf(t *testing.T) {
	tests := []struct {
		name   string
		s      []int
		target int
		want   int
	}{
		{"found", []int{1, 2, 3}, 2, 1},
		{"not found", []int{1, 2, 3}, 5, -1},
		{"first element", []int{1, 2, 3}, 1, 0},
		{"last element", []int{1, 2, 3}, 3, 2},
		{"empty slice", []int{}, 1, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IndexOf(tt.s, tt.target)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCopy(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		want []int
	}{
		{"basic", []int{1, 2, 3}, []int{1, 2, 3}},
		{"single", []int{42}, []int{42}},
		{"nil", nil, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Copy(tt.s)
			assert.Equal(t, tt.want, got)
			if tt.s != nil && len(tt.s) > 0 {
				tt.s[0] = 999
				assert.NotEqual(t, tt.s[0], got[0])
			}
		})
	}
}

func TestSort(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		want []int
	}{
		{"unsorted", []int{3, 1, 2}, []int{1, 2, 3}},
		{"sorted", []int{1, 2, 3}, []int{1, 2, 3}},
		{"reverse", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"single", []int{1}, []int{1}},
		{"empty", []int{}, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := make([]int, len(tt.s))
			copy(s, tt.s)
			Sort(s)
			assert.Equal(t, tt.want, s)
		})
	}
}

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name   string
		s      []int
		target int
		want   int
	}{
		{"found middle", []int{1, 2, 3, 4, 5}, 3, 2},
		{"found first", []int{1, 2, 3}, 1, 0},
		{"found last", []int{1, 2, 3}, 3, 2},
		{"not found", []int{1, 2, 3}, 5, -1},
		{"empty", []int{}, 1, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BinarySearch(tt.s, tt.target)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		name   string
		s      []int
		want   int
		wantOk bool
	}{
		{"basic", []int{1, 5, 3}, 5, true},
		{"negative", []int{-1, -5, -3}, -1, true},
		{"single", []int{42}, 42, true},
		{"empty", []int{}, 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := Max(tt.s)
			assert.Equal(t, tt.wantOk, ok)
			if ok {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		name   string
		s      []int
		want   int
		wantOk bool
	}{
		{"basic", []int{1, 5, 3}, 1, true},
		{"negative", []int{-1, -5, -3}, -5, true},
		{"single", []int{42}, 42, true},
		{"empty", []int{}, 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := Min(tt.s)
			assert.Equal(t, tt.wantOk, ok)
			if ok {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestSubslice(t *testing.T) {
	tests := []struct {
		name  string
		s     []int
		start int
		end   int
		want  []int
	}{
		{"basic", []int{1, 2, 3, 4, 5}, 1, 4, []int{2, 3, 4}},
		{"full", []int{1, 2, 3}, 0, 3, []int{1, 2, 3}},
		{"empty result", []int{1, 2, 3}, 1, 1, []int{}},
		{"invalid start", []int{1, 2, 3}, -1, 2, nil},
		{"invalid end", []int{1, 2, 3}, 0, 5, nil},
		{"start > end", []int{1, 2, 3}, 2, 1, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Subslice(tt.s, tt.start, tt.end)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestEdgeCases(t *testing.T) {
	t.Run("large slice performance", func(t *testing.T) {
		// Teste com slice grande
		s := make([]int, 10000)
		for i := range s {
			s[i] = i
		}
		
		Sort(s)
		idx := BinarySearch(s, 5000)
		assert.Equal(t, 5000, idx)
	})

	t.Run("negative numbers", func(t *testing.T) {
		s := []int{-5, -2, -8, -1}
		Sort(s)
		assert.Equal(t, []int{-8, -5, -2, -1}, s)
		
		max, _ := Max(s)
		min, _ := Min(s)
		assert.Equal(t, -1, max)
		assert.Equal(t, -8, min)
	})

	t.Run("duplicates", func(t *testing.T) {
		s := []int{2, 2, 2, 2}
		idx := IndexOf(s, 2)
		assert.Equal(t, 0, idx)
		
		sum := Sum(s)
		assert.Equal(t, 8, sum)
	})
}
