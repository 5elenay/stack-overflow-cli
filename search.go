package main

import (
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/5elenay/ezcli"
	"github.com/fatih/color"
	"github.com/pkg/browser"
)

func Search(c *ezcli.Command) {
	if len(c.CommandData.Arguments) < 1 {
		log.Fatal("Error: Add an argument to search.")
	}

	// Escape Url
	query := url.QueryEscape(c.CommandData.Arguments[0])

	// Get JSON Body and Convert
	var itemList StackOverflowItemList
	err := RequestAndUpdate(fmt.Sprintf("https://api.stackexchange.com/search/advanced?site=stackoverflow.com&sort=votes&q=%s", query), &itemList)

	if err != nil {
		log.Fatalf("Unknown Error: %s", err)
	}

	// Loop the items
	for index, item := range itemList.Items {
		// Tons of shitcode for colorize
		color.Unset()
		fmt.Print("Question ")
		color.Set(color.FgCyan)
		fmt.Printf("#%d\n", index+1)
		color.Unset()
		fmt.Print("  Title: ")
		color.Set(color.FgBlue)
		fmt.Printf("\"%s\"\n", item.Title)
		color.Unset()
		fmt.Print("  Tags: [ ")

		for _, i := range item.Tags {
			color.Set(color.FgMagenta)
			fmt.Print(i)
			color.Unset()
			fmt.Print(", ")
		}
		fmt.Print("]\n  Visit the Page? [")
		color.Set(color.FgGreen)
		fmt.Print("y")
		color.Unset()
		fmt.Print("/")
		color.Set(color.FgRed)
		fmt.Print("n")
		color.Unset()

		question := ezcli.Question{
			Input: "]: ",
		}

		question.Ask()

		// Open the URL in browser if answer is y or yes
		if strings.EqualFold(question.Answer, "y") || strings.EqualFold(question.Answer, "yes") {
			browser.OpenURL(item.Link)
		}
	}
}
