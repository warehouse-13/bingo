package templates

import "fmt"

func GenerateApp(name string) []byte {
	content := `package command

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

	app.Name = "%s"
	app.Usage = "REPLACE_ME"
	app.EnableBashCompletion = true
	app.Commands = commands()

	return app
}

func commands() []*cli.Command {
	return []*cli.Command{
		exampleCommand(),
		versionCommand(),
	}
}
`
	return []byte(fmt.Sprintf(content, name))
}
