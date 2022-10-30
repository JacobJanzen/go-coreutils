package main

import (
	"fmt"
	"os"

	"github.com/JacobJanzen/go-coreutils/pkg/cmd_options"
	"github.com/JacobJanzen/go-coreutils/pkg/system"
	"github.com/JacobJanzen/go-coreutils/pkg/templates"
	"golang.org/x/sys/unix"
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
	opts := templates.BasicOpts
	opts.HelpMessage.Usage = "whoami [OPTION]..."
	opts.HelpMessage.Description = "Print the user name associated with the current effective user ID."
	_ = cmd_options.ParseOptionsFromArgs(os.Args[1:], &opts)

	if opts.CallBasicOptions() {
		return
	}

	uid := unix.Geteuid()
	username, err := system.GetUsername(uint(uid))
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

	fmt.Println(username)
}
