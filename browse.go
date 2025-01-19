package main

import (
	"fmt"
	"os"

	"github.com/carpeliam/gitshorty/browse"
)

func main() {
	apiToken := os.Getenv("SHORTCUT_API_TOKEN")
	shortcut := browse.NewShortcutReader(apiToken)
	browser := browse.NewBrowser()
	git := browse.NewRepository()
	error := browse.BrowseStory(git, shortcut, browser)
	if error != nil {
		fmt.Println(error)
		os.Exit(1)
	}
}
