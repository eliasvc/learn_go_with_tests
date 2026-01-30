package main

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d given %v, want %d", got, numbers, want)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{1, 2, 3, 4, 5})
	want := []int{6, 15}

	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func BenchmarkSumAll(b *testing.B) {
	for b.Loop() {
		SumAll([]int{1, 2, 3}, []int{1, 2, 3, 4, 5})
	}
}
