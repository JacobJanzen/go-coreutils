package cmd_options

type Options interface {
	SetShortOption(opt rune)
	SetLongOption(opt string)
}

func ParseOptionsFromArgs(args []string, opts Options) (realArgs []string) {
	for _, arg := range args {
		if arg[0] != '-' {
			realArgs = append(realArgs, arg)
		} else if len(arg) >= 2 && arg[1] == '-' {
			opts.SetLongOption(arg[2:])
		} else if len(arg) >= 1 {
			for _, char := range arg {
				if char != '-' {
					opts.SetShortOption(char)
				}
			}
		}
	}
	return
}
