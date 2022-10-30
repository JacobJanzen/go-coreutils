package main

import (
	"fmt"
	"os"

	"github.com/JacobJanzen/go-coreutils/pkg/cmd_options"
	"github.com/JacobJanzen/go-coreutils/pkg/templates"
)

func main() {
	basicOpts := templates.BasicOpts
	setBasicOpts(&basicOpts)

	opts := Option{
		basicOpts: &basicOpts,
	}

	_ = cmd_options.ParseOptionsFromArgs(os.Args[1:], &opts)

	if basicOpts.CallBasicOptions() {
		return
	}

	if opts.err != nil {
		fmt.Fprintf(os.Stderr, "%s", opts.err.Error())
		os.Exit(1)
	}

	if nproc, err := getNProc(&opts); err != nil {
		fmt.Fprintf(os.Stderr, "%s", opts.err.Error())
	} else {
		fmt.Println(nproc)
	}
}
