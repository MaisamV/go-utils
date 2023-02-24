package utils

import (
	"testing"
)

func TestContains(t *testing.T) {
	a := 1
	b := []int{1, 2, 3}
	expected := true
	if Contains(a, b) != expected {
		t.Errorf("Contains(%d, %d) result is not %t", a, b, expected)
	}

	a = -1
	b = []int{1, -2, 3}
	expected = false
	if Contains(a, b) != expected {
		t.Errorf("Contains(%d, %d) result is not %t", a, b, expected)
	}

	c := 1.2
	d := []float64{1.2, -2.1, 3.2}
	expected = true
	if Contains(c, d) != expected {
		t.Errorf("Contains(%f, %f) result is not %t", c, d, expected)
	}

	c = 1.2
	d = []float64{1.1999999, -2.1, 3.2}
	expected = false
	if Contains(c, d) != expected {
		t.Errorf("Contains(%f, %f) result is not %t", c, d, expected)
	}

	e := "Hi"
	f := []string{"hi", "H i", "Yo"}
	expected = false
	if Contains(e, f) != expected {
		t.Errorf("Contains(%s, %s) result is not %t", e, f, expected)
	}

	e = "Yo"
	f = []string{"Yo", "hi", "H i"}
	expected = true
	if Contains(e, f) != expected {
		t.Errorf("Contains(%s, %s) result is not %t", e, f, expected)
	}

	e = "Yo"
	f = []string{}
	expected = false
	if Contains(e, f) != expected {
		t.Errorf("Contains(%s, %s) result is not %t", e, f, expected)
	}

	e = "Yo"
	f = []string{}
	f = nil
	expected = false
	if Contains(e, f) != expected {
		t.Errorf("Contains(%s, %s) result is not %t", e, f, expected)
	}
}
