package templates

import (
	"fmt"

	"github.com/JacobJanzen/go-coreutils/pkg/basic_options"
)

var helpShort = 'h'
var helpLong = "help"
var helpDesc = "display this help and exit"

var versionShort = 'v'
var versionLong = "version"
var versionDesc = "output version information and exit"

var copyright = "Copyright (C) 2022 Jacob Janzen"
var license = fmt.Sprintf("License GPLv3+: GNU GPL version 3 or later <https://gnu.org/licenses/gpl.html>.\nThis is free software: you are free to change and redistribute it.\nThere is NO WARRANTY, to the extent permitted by law.")
var author = "Jacob Janzen"

var BasicOpts = basic_options.BasicOptions{
	HelpMessage: basic_options.Help{
		Usage:       "",
		Description: "",
		Options: []basic_options.Option{
			{
				Short:       &helpShort,
				Long:        &helpLong,
				Description: &helpDesc,
			},
			{
				Short:       &versionShort,
				Long:        &versionLong,
				Description: &versionDesc,
			},
		},
	},
	VersionMessage: basic_options.Version{
		Name:      "",
		Version:   "",
		Copyright: &copyright,
		License:   &license,
		Author:    &author,
	},
}
