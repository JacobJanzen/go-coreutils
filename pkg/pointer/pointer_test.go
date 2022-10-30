package pointer

import "testing"

func TestOfChar(t *testing.T) {
	char := Of('a')

	if char == nil {
		t.Errorf("unexpected null pointer")
	} else if *char != 'a' {
		t.Errorf("expected 'a', got '%c'", *char)
	}
}

func TestOfInt(t *testing.T) {
	num := Of(1)

	if num == nil {
		t.Errorf("unexpected null pointer")
	} else if *num != 1 {
		t.Errorf("expected '1', got '%d'", *num)
	}
}

func TestOfString(t *testing.T) {
	str := Of("test")

	if str == nil {
		t.Errorf("unexpected null pointer")
	} else if *str != "test" {
		t.Errorf("expected \"test\", got \"%s\"", *str)
	}
}
