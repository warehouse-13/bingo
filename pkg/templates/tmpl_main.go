package templates

func GenerateMain(_ string) []byte {
	content := `package main

import (
	"log"
	"os"

  // Remove the dot '.'
	// And change the placeholder to your module
	. "REPLACE_ME/pkg/command"
)

func main() {
	app := command.NewApp(os.Stdout)

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
`
	return []byte(content)
}
