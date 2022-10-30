package basic_options

import "fmt"

type Option struct {
	Short       *rune
	Long        *string
	Description *string
}

type Help struct {
	Usage       string
	Description string
	Options     []Option
}

type Version struct {
	Name      string
	Version   string
	Copyright *string
	License   *string
	Author    *string
}

type BasicOptions struct {
	version        bool
	help           bool
	HelpMessage    Help
	VersionMessage Version
}

func (o *BasicOptions) SetShortOption(opt rune) {
	switch opt {
	case 'v':
		o.version = true
	case 'h':
		o.help = true
	}
}
func (o *BasicOptions) SetLongOption(opt string) {
	switch opt {
	case "version":
		o.version = true
	case "help":
		o.help = true
	}
}

func (o BasicOptions) CallBasicOptions() bool {
	if o.help {
		fmt.Print(o.getHelpMessage())
	}
	if o.version {
		fmt.Print(o.getVersionMessage())
	}

	return o.version || o.help
}

func (o BasicOptions) getVersionMessage() string {
	versionMessage := fmt.Sprintf("%s %s\n",
		o.VersionMessage.Name,
		o.VersionMessage.Version,
	)

	if o.VersionMessage.Copyright != nil {
		versionMessage = fmt.Sprintf("%s%s\n",
			versionMessage,
			*o.VersionMessage.Copyright,
		)
	}

	if o.VersionMessage.License != nil {
		versionMessage = fmt.Sprintf("%s%s\n",
			versionMessage,
			*o.VersionMessage.License,
		)
	}

	if o.VersionMessage.Author != nil {
		versionMessage = fmt.Sprintf("%s\nWritten by %s.\n",
			versionMessage,
			*o.VersionMessage.Author,
		)
	}

	return versionMessage
}

func (o BasicOptions) getHelpMessage() string {
	helpMessage := fmt.Sprintf("Usage: %s\n%s\n\n",
		o.HelpMessage.Usage,
		o.HelpMessage.Description,
	)

	longOptionLen := 0
	for _, opt := range o.HelpMessage.Options {
		if opt.Long != nil && len([]rune(*opt.Long)) > longOptionLen {
			longOptionLen = len([]rune(*opt.Long))
		}
	}

	for _, opt := range o.HelpMessage.Options {
		helpMessage = fmt.Sprintf("%s  ", helpMessage)
		if opt.Short != nil {
			helpMessage = fmt.Sprintf("%s-%c", helpMessage, *opt.Short)
			if opt.Long != nil {
				helpMessage = fmt.Sprintf("%s, ", helpMessage)
			} else {
				helpMessage = fmt.Sprintf("%s  ", helpMessage)
			}
		} else {
			helpMessage = fmt.Sprintf("%s    ", helpMessage)
		}

		if opt.Long != nil {
			helpMessage = fmt.Sprintf("%s--%s", helpMessage, *opt.Long)
			for i := 0; i < longOptionLen+2-len([]rune(*opt.Long)); i++ {
				helpMessage = fmt.Sprintf("%s ", helpMessage)
			}
		} else {
			for i := 0; i < longOptionLen+4; i++ {
				helpMessage = fmt.Sprintf("%s ", helpMessage)
			}
		}

		if opt.Description != nil {
			helpMessage = fmt.Sprintf("%s%s", helpMessage, *opt.Description)
		}
		helpMessage = fmt.Sprintf("%s\n", helpMessage)
	}

	return helpMessage
}
