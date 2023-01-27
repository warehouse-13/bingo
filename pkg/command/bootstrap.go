package command

import (
	"fmt"
	"os"

	"github.com/spf13/afero"
	"github.com/urfave/cli/v2"
	"github.com/warehouse-13/bingo/pkg/config"
	"github.com/warehouse-13/bingo/pkg/flags"
	"github.com/warehouse-13/bingo/pkg/templates"
)

func bootstrapCommand() *cli.Command {
	cfg := &config.Config{}
	fs := afero.NewOsFs()

	return &cli.Command{
		Name:    "bootstrap",
		Usage:   "bootstrap a cli framework",
		Aliases: []string{"b"},
		Before:  flags.ParseFlags(cfg),
		Flags: flags.CLIFlags(
			flags.WithCliNameFlag(),
		),
		Action: func(c *cli.Context) error {
			return BootstrapFn(cfg, fs)
		},
	}
}

func BootstrapFn(cfg *config.Config, fs afero.Fs) error {
	for _, f := range templates.ManagedFiles() {
		if err := mustNotExist(fs, f.Name); err != nil {
			return err
		}
	}

	fmt.Printf("Bootstrapping CLI for `%s`...\n\n", cfg.CliName)

	for _, d := range templates.ManagedDirs() {
		ok, err := afero.DirExists(fs, d)
		if err != nil {
			return err
		}

		if !ok {
			if err := fs.MkdirAll(d, 0755); err != nil {
				return err
			}

			fmt.Printf("Created directory %s\n", d)
		}
	}

	// We check and write in separate loops, as we don't want to write some and
	// abort on others.
	for _, f := range templates.ManagedFiles() {
		if err := afero.WriteFile(fs, f.Name, f.Contents(cfg.CliName), 0644); err != nil {
			return err
		}

		fmt.Printf("Written file %s\n", f.Name)
	}

	fmt.Println("\n\nDone!")
	fmt.Println("Check each file and follow instructions within.")
	fmt.Println("Run `go mod tidy` and start using.")

	return nil
}

func mustNotExist(fs afero.Fs, path string) error {
	if _, err := fs.Stat(path); os.IsNotExist(err) {
		return nil
	}

	return fmt.Errorf("expected file `%s` to not exist, aborting", path)
}
