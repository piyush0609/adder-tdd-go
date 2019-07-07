package integers

import (
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("expects seven when three and four are added", func(t *testing.T) {
		sum := Add(3, 4)
		expected := 7

		if sum != expected {
			t.Errorf("Expected '%d' got '%d'", expected, sum)
		}
	})

	t.Run("expects eight when four and four are added", func(t *testing.T) {
		sum := Add(4, 4)
		expected := 8

		if sum != expected {
			t.Errorf("Expected '%d' got '%d'", expected, sum)
		}
	})
}
