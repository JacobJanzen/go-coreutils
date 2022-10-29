package main

import (
	"fmt"
	"os"

	/*
	   #include <pwd.h>
	   #include <sys/types.h>
	*/
	"C"

	"github.com/JacobJanzen/go-coreutils/pkg/cmd_options"
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
	opts := Options{}
	_ = cmd_options.ParseOptionsFromArgs(os.Args[1:], &opts)

	if opts.help {
		printHelpMessage()
		return
	}
	if opts.version {
		printVersionMessage()
		return
	}

	uid := unix.Geteuid()
	pw := C.getpwuid(C.uint(uid))

	if pw != nil {
		fmt.Println(C.GoString(pw.pw_name))
	}
}
