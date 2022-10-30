package main

import (
	"fmt"
	"os"

	"github.com/JacobJanzen/go-coreutils/pkg/cmd_options"
	"github.com/JacobJanzen/go-coreutils/pkg/system"
	"github.com/JacobJanzen/go-coreutils/pkg/templates"
	"golang.org/x/sys/unix"
)

func main() {
	opts := templates.BasicOpts
	opts.HelpMessage.Usage = "whoami [OPTION]..."
	opts.HelpMessage.Description = "Print the user name associated with the current effective user ID."
	_ = cmd_options.ParseOptionsFromArgs(os.Args[1:], &opts)

	if opts.CallBasicOptions() {
		return
	}

	uid := unix.Geteuid()
	if uid < 0 {
		fmt.Fprintf(os.Stderr, "error invalid user id\n")
	}

	username, err := system.GetUsername(uint(uid))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}

	fmt.Println(username)
}
