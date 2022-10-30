package main

import (
	"fmt"
	"strconv"

	"github.com/JacobJanzen/go-coreutils/pkg/basic_options"
	"github.com/JacobJanzen/go-coreutils/pkg/pointer"
	"github.com/JacobJanzen/go-coreutils/pkg/system"
)

type Option struct {
	basicOpts *basic_options.BasicOptions
	all       bool
	ignore    int
	err       error
}

func getNProc(o *Option) (int, error) {
	var nproc int
	var err error
	if o.all {
		nproc, err = system.NProcAll()
	} else {
		nproc, err = system.NProc()
	}

	if err != nil {
		return nproc, fmt.Errorf("error getting number of processors: %v", err)
	}

	if nproc > o.ignore {
		nproc -= o.ignore
	} else {
		nproc = 1
	}

	return nproc, nil
}

func setBasicOpts(basicOpts *basic_options.BasicOptions) {
	basicOpts.HelpMessage.Usage = "nproc [OPTION]..."
	basicOpts.HelpMessage.Description = `Print the number of processing units available to the current process,
which may be less than the number of online processors`
	basicOpts.HelpMessage.Options = append(
		basicOpts.HelpMessage.Options,
		basic_options.Option{
			Short:       pointer.Of('a'),
			Long:        pointer.Of("all"),
			Description: pointer.Of("print the number of installed processors"),
		},
	)
	basicOpts.HelpMessage.Options = append(
		basicOpts.HelpMessage.Options,
		basic_options.Option{
			Long:        pointer.Of("ignore=N"),
			Description: pointer.Of("if possible, exclude N processing units"),
		},
	)
	basicOpts.VersionMessage.Name = "nproc"
	basicOpts.VersionMessage.Version = "1.0"
}

func (o *Option) SetShortOption(opt rune) {
	switch opt {
	case 'a':
		o.all = true
	default:
		o.basicOpts.SetShortOption(opt)
	}
}

func (o *Option) SetLongOption(opt string) {
	if opt == "all" {
		o.all = true
	} else if len(opt) >= 7 && opt[:7] == "ignore=" {
		num, err := strconv.Atoi(opt[7:])
		if err != nil {
			o.err = fmt.Errorf("invalid number: '%v'", opt[:7])
		} else {
			o.ignore = num
		}
	} else {
		o.basicOpts.SetLongOption(opt)
	}
}
