package main

import (
	"fmt"
	"os"

	"github.com/JacobJanzen/go-coreutils/pkg/cmd_options"
	"github.com/JacobJanzen/go-coreutils/pkg/system"
	"github.com/JacobJanzen/go-coreutils/pkg/templates"
)

func main() {
	basicOpts := templates.BasicOpts
	setBasicOpts(&basicOpts)
	opts := Options{
		basicOpts: &basicOpts,
	}

	_ = cmd_options.ParseOptionsFromArgs(os.Args[1:], &opts)

	if opts.basicOpts.CallBasicOptions() {
		return
	}

	tty, err := system.GetTTY()
	if err != nil {
		if !opts.silent {
			fmt.Fprintln(os.Stderr, err)
		}
		os.Exit(1)
	} else if !opts.silent {
		fmt.Println(tty)
	}
}
