package main

import (
	"github.com/JacobJanzen/go-coreutils/pkg/basic_options"
	"github.com/JacobJanzen/go-coreutils/pkg/pointer"
)

type Options struct {
	silent    bool
	basicOpts *basic_options.BasicOptions
}

func setBasicOpts(basicOpts *basic_options.BasicOptions) {
	basicOpts.HelpMessage.Usage = "tty [OPTION]..."
	basicOpts.HelpMessage.Description = "Print the file name of the terminal connected to standard input."
	basicOpts.VersionMessage.Name = "yes"
	basicOpts.VersionMessage.Version = "1.0"

	basicOpts.HelpMessage.Options = append(
		basicOpts.HelpMessage.Options,
		basic_options.Option{
			Short: pointer.Of('s'),
			Long:  pointer.Of("silent"),
		},
	)

	basicOpts.HelpMessage.Options = append(
		basicOpts.HelpMessage.Options,
		basic_options.Option{
			Long:        pointer.Of("quiet"),
			Description: pointer.Of("print nothing, only return an exit status"),
		},
	)
}

func (o *Options) SetShortOption(opt rune) {
	switch opt {
	case 's':
		o.silent = true
	default:
		o.basicOpts.SetShortOption(opt)
	}
}

func (o *Options) SetLongOption(opt string) {
	switch opt {
	case "silent":
		o.silent = true
	case "quiet":
		o.silent = true
	default:
		o.basicOpts.SetLongOption(opt)
	}
}
