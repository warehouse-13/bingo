package main

import (
	"fmt"
	"os"

	"github.com/warehouse-13/bingo/pkg/command"
)

func main() {
	app := command.NewApp(os.Stdout)

	if err := app.Run(os.Args); err != nil {
		fmt.Println("\n", err)
		os.Exit(1)
	}
}
