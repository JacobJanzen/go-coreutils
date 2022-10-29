package cmd_options

import "testing"

type TestOption struct {
	OptA bool
	OptB bool
	OptC bool
	OptD bool
	OptE bool
}

func (o *TestOption) SetShortOption(opt rune) {
	switch opt {
	case 'a':
		o.OptA = true
	case 'b':
		o.OptB = true
	case 'c':
		o.OptC = true
	case 'd':
		o.OptD = true
	case 'e':
		o.OptE = true
	}
}

func (o *TestOption) SetLongOption(opt string) {
	switch opt {
	case "OPTIONA":
		o.OptA = true
	case "OPTIONB":
		o.OptB = true
	case "OPTIONC":
		o.OptC = true
	case "OPTIOND":
		o.OptD = true
	case "OPTIONE":
		o.OptE = true
	}
}

func TestParseOptionsFromArgsRemovesOptions(t *testing.T) {
	args := []string{
		"test1",
		"test2",
		"--test3",
		"-test4",
		"test5",
	}
	expectedArgs := []string{
		"test1",
		"test2",
		"test5",
	}
	opts := TestOption{}
	realArgs := ParseOptionsFromArgs(args, &opts)

	if len(realArgs) != len(expectedArgs) {
		t.Errorf("expected arg list of length %d, got list of length %d", len(expectedArgs), len(realArgs))
	}
}

func TestParseOptionsFromArgsSetsOptions(t *testing.T) {
	args := []string{
		"-ab",
		"--OPTIONC",
		"--OPTIOND",
	}
	expectedOptions := TestOption{
		OptA: true,
		OptB: true,
		OptC: true,
		OptD: true,
		OptE: false,
	}
	opts := TestOption{}
	_ = ParseOptionsFromArgs(args, &opts)

	if opts != expectedOptions {
		t.Errorf("expected %v, got %v", expectedOptions, opts)
	}
}
