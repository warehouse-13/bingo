package templates

func GenerateFlags(_ string) []byte {
	content := `package flags

import (
	"github.com/urfave/cli/v2"

  // Remove the dot '.'
	// And change the placeholder to your module
	. "REPLACE_ME/pkg/config"
)

// WithFlagsFunc can be used with CLIFlags to build a list of flags for a
// command.
type WithFlagsFunc func() []cli.Flag

// CLIFlags takes a list of WithFlagsFunc options and returns a list of flags
// for a command.
func CLIFlags(options ...WithFlagsFunc) []cli.Flag {
	flags := []cli.Flag{}

	for _, group := range options {
		flags = append(flags, group()...)
	}

	return flags
}

const exampleFlag = "name"

// WithExampleFlag adds an example flag to the command
func WithExampleFlag() WithFlagsFunc {
	return func() []cli.Flag {
		return []cli.Flag{
			&cli.StringFlag{
				Name:     exampleFlag,
				Aliases:  []string{"n"},
				Usage:    "an example flag",
				Required: true,
			},
		}
	}
}

// ParseFlags processes all flags on the CLI context and builds a config object
// which will be used in the command's action.
func ParseFlags(cfg *config.Config) cli.BeforeFunc {
	return func(ctx *cli.Context) error {
		cfg.ExampleOpt = ctx.String(exampleFlag)

		return nil
	}
}
`
	return []byte(content)
}
