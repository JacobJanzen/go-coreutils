package encoding

import (
	"reflect"
	"testing"
)

func TestMap5BitBlockToValue(t *testing.T) {
	a := Map5BitBlockToValue(0)
	f := Map5BitBlockToValue(5)
	two := Map5BitBlockToValue(26)
	six := Map5BitBlockToValue(30)
	if a != 'A' {
		t.Errorf("expected 'A', got '%c'", a)
	}
	if f != 'F' {
		t.Errorf("expected 'F', got '%c'", f)
	}
	if two != '2' {
		t.Errorf("expected '2', got '%c'", two)
	}
	if six != '6' {
		t.Errorf("expected '6', got '%c'", six)
	}
}

func TestMap5BitBlockToValueTooManyBits(t *testing.T) {
	var in byte = 0b100000
	in += 5
	f := Map5BitBlockToValue(in)

	if f != 'F' {
		t.Errorf("expected 'F', got '%c'", f)
	}
}

func TestMap5ByteBlockInto5BitGroups(t *testing.T) {
	expected := [8]byte{
		0b00000,
		0b11111,
		0b00000,
		0b11111,
		0b00000,
		0b11111,
		0b00000,
		0b11111,
	}

	input := [5]byte{
		0b00000111,
		0b11000001,
		0b11110000,
		0b01111100,
		0b00011111,
	}

	result := Map5ByteBlockInto5BitGroups(input[:])

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected\n%v\ngot\n%v\n", expected, result)
	}
}

func TestMap5ByteBlockInto5BitGroupsLessThan5BytesInput(t *testing.T) {
	expected := [8]byte{
		0b00000,
		0b11111,
		0b00000,
		0b11111,
		0b00000,
		0b11111,
		0b00000,
		0b00000,
	}

	input := [4]byte{
		0b00000111,
		0b11000001,
		0b11110000,
		0b01111100,
	}

	result := Map5ByteBlockInto5BitGroups(input[:])

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected\n%v\ngot\n%v\n", expected, result)
	}
}
