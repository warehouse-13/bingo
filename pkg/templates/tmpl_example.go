package templates

func GenerateExample(_ string) []byte {
	content := `package command

import (
	"fmt"

	"github.com/urfave/cli/v2"

  // Remove the dots '.'
	// And change the placeholders to your module
	. "REPLACE_ME/pkg/config"
	. "REPLACE_ME/pkg/flags"
)

func exampleCommand() *cli.Command {
	cfg := &config.Config{}

	return &cli.Command{
		Name:    "example",
		Usage:   "an example command",
		Aliases: []string{"e"},
		Before:  flags.ParseFlags(cfg),
		Flags: flags.CLIFlags(
			flags.WithExampleFlag(),
		),
		Action: func(c *cli.Context) error {
			return ExampleFn(cfg)
		},
	}
}

func ExampleFn(cfg *config.Config) error {
	fmt.Println(cfg.ExampleOpt)

	return nil
}
`
	return []byte(content)
}
