package integers

import "testing"

func TestAdder(t *testing.T) {
	assertInteger := func(t *testing.T, expected, want int) {
		t.Helper()
		if want != expected {
			t.Errorf("Expcted '%d' but got '%d'", expected, want)
		}
	}

	assertInteger(t, 4, Add(2, 2))
	assertInteger(t, 12, Add(5, 7))
}
