package arrays

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("sum of slice", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("sum of all slices", func(t *testing.T) {
		p := []int{1, 2}
		q := []int{0, 9}

		got := SumAll(p, q)
		want := []int{3, 9}

		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v, given %v %v", got, want, p, q)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int, inputs ...[]int) {
		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v, given %v", got, want, inputs)
		}
	}

	t.Run("sum of all tails of slices", func(t *testing.T) {
		p := []int{1, 2, 3}
		q := []int{0, 1, 9}

		got := SumAllTails(p, q)
		want := []int{5, 10}

		checkSums(t, got, want, p, q)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		p := []int{}
		q := []int{0, 1, 9}

		got := SumAllTails(p, q)
		want := []int{0, 10}

		checkSums(t, got, want, p, q)
	})
}
