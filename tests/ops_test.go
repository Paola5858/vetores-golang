package tests

import (
	"testing"

	"github.com/Paola5858/vetores-golang/slices"
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
			got := slices.AppendSafe(tt.s, tt.val)
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
			got := slices.Reverse(tt.s)
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
			got := slices.IndexOf(tt.s, tt.target)
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
		{"empty", []int{}, []int{}},
		{"nil", nil, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := slices.Copy(tt.s)
			assert.Equal(t, tt.want, got)
			if tt.s != nil && len(tt.s) > 0 {
				// Verify it's a different slice (not same reference)
				tt.s[0] = 999
				assert.NotEqual(t, tt.s[0], got[0])
			}
		})
	}
}

