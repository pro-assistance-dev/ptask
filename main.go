package main

import (
	"ptask/cmd/cli"
	"ptask/cmd/works"
	"log"
)

var ScriptsPath string

func main() {
	if ScriptsPath == "" {
		log.Fatal("scripts path is required")
	}
	Select()
}

func Select() {
	works.ParseWorks(ScriptsPath)
	worksList := works.GetWorks()
	selected := cli.Select(worksList)

	if selected == nil {
		return
	}
	for _, work := range selected {
		work.Open(ScriptsPath)
	}
}
