package arraynslices

// go test -cover

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		array_numbers := []int{1, 2, 3, 4, 5}

		got := Sum(array_numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, array_numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		slice_numbers := []int{1, 2, 3}

		got := Sum(slice_numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, slice_numbers)
		}
	})
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) { //convenient way of comparing slices (and other things) you must be careful when using it
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of tails of", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})

}
