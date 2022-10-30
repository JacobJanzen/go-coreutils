package main

import (
	"fmt"
	"os"

	"github.com/JacobJanzen/go-coreutils/pkg/cmd_options"
	"github.com/JacobJanzen/go-coreutils/pkg/templates"
)

func main() {
	opts := templates.BasicOpts
	opts.HelpMessage.Usage = "yes [STRING]..."
	opts.HelpMessage.Description = "Repeatedly output a line with all specified STRING(s), or 'y'."
	opts.VersionMessage.Name = "yes"
	opts.VersionMessage.Version = "1.0"

	args := os.Args[1:]
	args = cmd_options.ParseOptionsFromArgs(args, &opts)

	if opts.CallBasicOptions() {
		return
	}

	printY := true
	if len(args) > 0 {
		printY = false
	}

	for {
		if printY {
			fmt.Println("y")
		} else {
			for _, word := range args {
				fmt.Printf("%s ", word)
			}
			fmt.Println()
		}
	}
}
