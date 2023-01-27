package flags

import (
	"github.com/urfave/cli/v2"

	"github.com/warehouse-13/bingo/pkg/config"
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

const nameFlag = "cli-name"

// WithCliNameFlag adds the flag to the command
func WithCliNameFlag() WithFlagsFunc {
	return func() []cli.Flag {
		return []cli.Flag{
			&cli.StringFlag{
				Name:     nameFlag,
				Aliases:  []string{"n"},
				Usage:    "the name of the cli tool to bootstrap (required)",
				Required: true,
			},
		}
	}
}

// ParseFlags processes all flags on the CLI context and builds a config object
// which will be used in the command's action.
func ParseFlags(cfg *config.Config) cli.BeforeFunc {
	return func(ctx *cli.Context) error {
		cfg.CliName = ctx.String(nameFlag)

		return nil
	}
}
