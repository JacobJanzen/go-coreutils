package main

import (
	"fmt"
	"os"

	"github.com/JacobJanzen/go-coreutils/pkg/cmd_options"
)

type Options struct {
	version bool
	help    bool
}

func (o *Options) SetShortOption(opt rune) {}
func (o *Options) SetLongOption(opt string) {
	switch opt {
	case "version":
		o.version = true
	case "help":
		o.help = true
	}
}

func printHelpMessage() {
	fmt.Println("HELP MESSAGE")
}

func printVersionMessage() {
	fmt.Println("VERSION MESSAGE")
}

func main() {
	opts := Options{}
	args := os.Args[1:]
	args = cmd_options.ParseOptionsFromArgs(args, &opts)

	if opts.help {
		printHelpMessage()
		return
	}
	if opts.version {
		printVersionMessage()
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
