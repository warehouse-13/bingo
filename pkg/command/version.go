package command

import (
	"encoding/json"
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/warehouse-13/bingo/pkg/version"
)

func versionCommand() *cli.Command {
	return &cli.Command{
		Name:    "version",
		Usage:   "print the version number for bingo",
		Aliases: []string{"v"},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "long",
				Value:   false,
				Aliases: []string{"l"},
				Usage:   "print the long version information",
			},
		},
		Action: VersionFn,
	}
}

type versionInfo struct {
	PackageName string
	Version     string
	CommitHash  string
	BuildDate   string
}

func VersionFn(ctx *cli.Context) error {
	if !ctx.Bool("long") {
		fmt.Println(version.Version)

		return nil
	}

	info := versionInfo{
		version.PackageName,
		version.Version,
		version.CommitHash,
		version.BuildDate,
	}

	out, err := json.MarshalIndent(info, "", "  ")
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", string(out))

	return nil
}
