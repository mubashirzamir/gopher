package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("non-zero repititons", func(t *testing.T) {
		repeated := Repeat("a", 4)
		expected := "aaaa"

		assertCorrectRepitions(t, repeated, expected)
	})

	t.Run("zero repititions", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := ""

		assertCorrectRepitions(t, repeated, expected)
	})

	t.Run("empty string, multiple repititions", func(t *testing.T) {
		repeated := Repeat("", 5)
		expected := ""

		assertCorrectRepitions(t, repeated, expected)
	})
}

func assertCorrectRepitions(t testing.TB, repeated, expected string) {
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	// setup
	for b.Loop() {
		// code to measure
		Repeat("a", 5)
	}
	// cleanup
}
