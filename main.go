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
			Usages:      []string{"search \"Python\"", "search \"firebase\" --sort=votes"},
			Options: []*ezcli.CommandOption{
				{
					Name:        "sort",
					Description: "Edit the sort order.\n      Avaible: activity, creation, votes, hot, week, month\n      Default: activity",
					Aliases:     []string{"order"},
				},
			},
			Execute: Search,
		},
	})

	app.Handle()
}
