package main

import (
	"github.com/5elenay/ezcli"
)

func main() {
	app := ezcli.NewApp("stovc")

	app.AddCommands([]*ezcli.Command{
		{
			Name:        "search",
			Aliases:     []string{"find", "s"},
			Description: "Search for a question in stack overflow.",
			Usages:      []string{"search \"Python\""},
			Execute:     Search,
		},
	})

	app.Handle()
}
