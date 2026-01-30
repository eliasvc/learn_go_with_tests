package main

import (
	"reflect"
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

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("sum tails of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{1, 2, 3, 4, 5})
		want := []int{5, 14}

		checkSums(t, got, want)
	})

	t.Run("safely sum when empty slices are used", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1, 2, 3})
		want := []int{0, 5}

		checkSums(t, got, want)
	})
}

func BenchmarkSumAll(b *testing.B) {
	for b.Loop() {
		SumAllTails([]int{1, 2, 3}, []int{1, 2, 3, 4, 5})
	}
}
