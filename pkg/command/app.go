package command

import (
	"io"

	"github.com/urfave/cli/v2"
)

// NewApp is a builder which returns a cli.App.
func NewApp(out io.Writer) *cli.App {
	app := cli.NewApp()

	if out != nil {
		app.Writer = out
	}

	app.Name = "bingo"
	app.Usage = "a basic go cli bootstrapper"
	app.EnableBashCompletion = true
	app.Commands = commands()

	return app
}

func commands() []*cli.Command {
	return []*cli.Command{
		bootstrapCommand(),
		versionCommand(),
	}
}
