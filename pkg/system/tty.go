package system

import (
	"fmt"
	"os"
)

func GetTTY() (string, error) {
	stdin, err := os.Readlink(os.Stdin.Name())
	if err != nil {
		return "", fmt.Errorf("error getting tty: failed to get stdin")
	}

	tty, err := os.Readlink(stdin)
	if err != nil {
		return "", fmt.Errorf("error getting tty: not a tty")
	}

	return tty, nil
}
