package basic_options

import "testing"

var short = 'a'
var long = "Test"
var longer = "Test2"
var description = "This is a description"
var copyright = "Copyright (C) 2022 Jacob Janzen"
var license = "License"
var author = "Jacob Janzen"

var opt = BasicOptions{
	HelpMessage: Help{
		Usage:       "test [options]",
		Description: "this is a test help message.",
		Options: []Option{
			{
				Short: &short,
			},
			{
				Long: &long,
			},
			{
				Long: &longer,
			},
			{
				Description: &description,
			},
			{
				Short:       &short,
				Description: &description,
			},
			{
				Long:        &long,
				Description: &description,
			},
			{
				Short:       &short,
				Long:        &long,
				Description: &description,
			},
		},
	},
	VersionMessage: Version{
		Name:      "test",
		Version:   "1.0",
		Copyright: &copyright,
		License:   &license,
		Author:    &author,
	},
}

const expected_help = `Usage: test [options]
this is a test help message.

  -a           
      --Test   
      --Test2  
               This is a description
  -a           This is a description
      --Test   This is a description
  -a, --Test   This is a description
`

const expected_version = `test 1.0
Copyright (C) 2022 Jacob Janzen
License

Written by Jacob Janzen.
`

func TestGetHelpMessage(t *testing.T) {
	help := opt.getHelpMessage()
	if help != expected_help {
		t.Errorf("Expected\n%s\nGot\n%s\n", expected_help, help)
	}
}

func TestGetVersionMessage(t *testing.T) {
	version := opt.getVersionMessage()
	if version != expected_version {
		t.Errorf("Expected\n%s\nGot%s\n", expected_version, version)
	}
}
