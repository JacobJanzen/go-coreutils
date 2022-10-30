package main

import (
	"fmt"
	"os"

	"github.com/JacobJanzen/go-coreutils/pkg/cmd_options"
	"github.com/JacobJanzen/go-coreutils/pkg/system"
	"github.com/JacobJanzen/go-coreutils/pkg/templates"
)

func main() {
	opts := templates.BasicOpts
	opts.HelpMessage.Usage = "logname [OPTION]..."
	opts.HelpMessage.Description = "Print the user's login name."
	opts.VersionMessage.Name = "logname"
	opts.VersionMessage.Version = "1.0"
	_ = cmd_options.ParseOptionsFromArgs(os.Args[1:], &opts)

	if opts.CallBasicOptions() {
		return
	}

	username, err := system.GetLogin()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}

	fmt.Println(username)
}
